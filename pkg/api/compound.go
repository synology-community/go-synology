package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
)

// CompoundElement is one sub-request wrapped inside a SYNO.Entry.Request
// compound envelope (see PLAT-705). Its wire encoding is FLAT: Params'
// fields merge directly onto the same JSON object as api/method/version --
// DSM's compound endpoint has no nested "params" wrapper, so MarshalJSON
// deliberately does not produce one.
//
// Some DSM endpoints -- SYNO.Core.Share.Permission set among them -- return
// 403 when called directly, with or without parameters, and only accept the
// call wrapped in this envelope. list-style endpoints generally work either
// way; wrap them too when a caller wants to batch several calls into one
// round trip.
type CompoundElement struct {
	API     string
	Method  string
	Version int

	// Params marshals to a JSON object whose fields merge onto this
	// element alongside api/method/version. A nil Params sends only those
	// three fields. A value that does not marshal to a JSON object (a
	// slice, a scalar, ...) is an error -- there is nowhere flat to put it.
	Params any
}

// MarshalJSON implements the flat merge described on CompoundElement.
func (e CompoundElement) MarshalJSON() ([]byte, error) {
	body := map[string]json.RawMessage{}
	if e.Params != nil {
		pb, err := json.Marshal(e.Params)
		if err != nil {
			return nil, fmt.Errorf(
				"compound element %s.%s: marshal params: %w",
				e.API, e.Method, err,
			)
		}
		if err := json.Unmarshal(pb, &body); err != nil {
			return nil, fmt.Errorf(
				"compound element %s.%s: params must marshal to a JSON object, got %s: %w",
				e.API, e.Method, pb, err,
			)
		}
	}

	apiJSON, err := json.Marshal(e.API)
	if err != nil {
		return nil, err
	}
	methodJSON, err := json.Marshal(e.Method)
	if err != nil {
		return nil, err
	}
	versionJSON, err := json.Marshal(e.Version)
	if err != nil {
		return nil, err
	}
	body["api"] = apiJSON
	body["method"] = methodJSON
	body["version"] = versionJSON

	return json.Marshal(body)
}

// CompoundRequest is the SYNO.Entry.Request envelope body. Compound marshals
// as a single JSON array via the "json" query-encoding option (see
// pkg/query), so each element's flat shape reaches the wire untouched -- no
// second layer of escaping.
type CompoundRequest struct {
	Compound      []CompoundElement `url:"compound,json"`
	Mode          string            `url:"mode"`
	StopWhenError bool              `url:"stop_when_error"`
}

// CompoundResult is one sub-request's outcome inside a compound response.
type CompoundResult struct {
	API     string          `json:"api"`
	Method  string          `json:"method"`
	Version int             `json:"version"`
	Success bool            `json:"success"`
	Data    json.RawMessage `json:"data,omitempty"`
	Error   ApiError        `json:"error"`
}

// CompoundResponse is the SYNO.Entry.Request envelope's data payload.
type CompoundResponse struct {
	HasFail bool             `json:"has_fail"`
	Result  []CompoundResult `json:"result"`
}

// CompoundFailedError reports that one sub-request inside a compound
// envelope failed. See CompoundError.
type CompoundFailedError struct {
	Index  int
	API    string
	Method string
	Err    error
}

func (e CompoundFailedError) Error() string {
	return fmt.Sprintf(
		"compound sub-request %d (%s.%s) failed: %v",
		e.Index, e.API, e.Method, e.Err,
	)
}

func (e CompoundFailedError) Unwrap() error { return e.Err }

// PostCompound sends one or more sub-requests wrapped in a SYNO.Entry.Request
// envelope. An empty mode defaults to "sequential" (both "sequential" and
// "parallel" are accepted by DSM). It returns an error only for an
// envelope-level failure (handled by Post/handle, same as any other call) or
// a transport error -- a partial failure recorded *inside* the envelope
// decodes successfully and does NOT surface here.
//
// Callers MUST pass the result through CompoundError (or inspect
// HasFail/Result themselves) before treating the call as having succeeded;
// see CompoundError's own warning and PLAT-705.
func PostCompound(
	c Api,
	ctx context.Context,
	elements []CompoundElement,
	mode string,
	stopWhenError bool,
) (*CompoundResponse, error) {
	if mode == "" {
		mode = "sequential"
	}
	req := CompoundRequest{
		Compound:      elements,
		Mode:          mode,
		StopWhenError: stopWhenError,
	}
	resp, err := Post[CompoundResponse](c, ctx, &req, Compound)
	if err != nil {
		return nil, err
	}

	// A compound that came back with fewer results than sub-requests sent has
	// not done what was asked, however cheerful the envelope looks. DSM reports
	// envelope-level success:true for such a response, so without this check a
	// caller doing `if err != nil` treats "nothing executed" as "everything
	// worked" -- the same silent-pass shape as PLAT-698 (an ACL checker that
	// reported clean when its API call had failed), PLAT-711 (a lint gate blind
	// to unstaged files) and PLAT-715 (a local-CI gate that skipped every job
	// and exited 0). Refusing to call an empty or short result set a success is
	// the whole point of those tickets, and it would be perverse to reintroduce
	// it here.
	if resp != nil && len(resp.Result) != len(elements) {
		return nil, fmt.Errorf(
			"compound returned %d result(s) for %d sub-request(s): the envelope "+
				"reported success but did not execute what was sent",
			len(resp.Result), len(elements),
		)
	}

	return resp, nil
}

// CompoundError reports the first sub-request in resp that failed, wrapping
// its DSM error as a CompoundFailedError. It returns nil only when every
// sub-request succeeded.
//
// THIS CHECK IS NOT OPTIONAL. The envelope can report success:true --
// PostCompound sees no error -- while a sub-request wrapped inside it
// failed: has_fail:true, or an individual result[i].success:false. A caller
// that only checks the error PostCompound returns will report success for a
// call that never actually happened. That exact defect class -- a check
// that reports success without having verified anything -- closed three
// tickets in one day (PLAT-698, PLAT-711, PLAT-715). Pass every
// PostCompound result through this before treating it as a success.
func CompoundError(resp *CompoundResponse) error {
	if resp == nil {
		return nil
	}
	for i, r := range resp.Result {
		if r.Success {
			continue
		}
		var err error
		if r.Error.Code == 0 {
			err = errors.New("sub-request failed with no error code reported")
		} else {
			err = r.Error.WithSummaries(GlobalErrors)
		}
		return CompoundFailedError{Index: i, API: r.API, Method: r.Method, Err: err}
	}
	if resp.HasFail {
		return errors.New(
			"compound envelope reported has_fail=true but no sub-request result reported failure",
		)
	}
	return nil
}
