package api

import (
	"encoding/json"
	"fmt"
	"maps"
	"strings"
)

// GlobalErrors holds mapping of global errors not related to particular API endpoint.
var GlobalErrors ErrorSummary = ErrorSummary{
	100:  "Unknown error",
	101:  "No parameter of API, method or version",
	102:  "The requested API does not exist",
	103:  "The requested method does not exist",
	104:  "The requested version does not support the functionality",
	105:  "The logged in session does not have permission",
	106:  "Session timeout",
	107:  "Session interrupted by duplicate login",
	119:  "SID not found",
	400:  "Invalid parameter of file operation",
	401:  "Unknown error of file operation",
	402:  "System is too busy",
	403:  "Invalid user does this file operation",
	404:  "Invalid group does this file operation",
	405:  "Invalid user and group does this file operation",
	406:  "Can't get user/group information from the account server",
	407:  "Operation not permitted",
	408:  "No such file or directory",
	409:  "Non-supported file system",
	410:  "Failed to connect internet-based file system (e.g., CIFS)",
	411:  "Read-only file system",
	412:  "Filename too long in the non-encrypted file system",
	413:  "Filename too long in the encrypted file system",
	414:  "File already exists",
	415:  "Disk quota exceeded",
	416:  "No space left on device",
	417:  "Input/output error",
	418:  "Illegal name or path",
	419:  "Illegal file name",
	420:  "Illegal file name on FAT file system",
	421:  "Device or resource busy",
	599:  "No such task of the file operation",
	1400: "Failed to extract files.",
	1401: "Cannot open the file as archive.",
	1402: "Failed to read archive data error",
	1403: "Wrong password.",
	1404: "Failed to get the file and dir list in an archive.",
	1405: "Failed to find the item ID in an archive file.",
	1800: "There is no Content-Length information in the HTTP header or the received size doesn't match the value of Content-Length information in the HTTP header.",
	1801: "Wait too long, no date can be received from client (Default maximum wait time is 3600 seconds).",
	1802: "No filename information in the last part of file content.",
	1803: "Upload connection is cancelled.",
	1804: "Failed to upload oversized file to FAT file system.",
	1805: "Can't overwrite or skip the existing file, if no   parameter is given.",
}

func (es ErrorSummary) Combine(params ...ErrorSummary) ErrorSummary {
	for _, p := range params {
		maps.Copy(es, p)
	}

	return es
}

// ErrorDescriber defines interface to report all known errors to particular object.
type ErrorDescriber interface {
	// ErrorSummaries returns information about all known errors.
	ErrorSummaries() []ErrorSummary
}

// ApiError defines a structure for error object returned by Synology API.
// It is a high-level error for a particular API family.
type ApiError struct {
	Code    int         `json:"code,omitempty"`
	Summary string      `json:"-"`
	Errors  []ErrorItem `json:"-"`
}

// ErrorItem defines detailed request error.
type ErrorItem struct {
	Code    int
	Summary string
	Details ErrorFields
}

// ErrorSummary is a simple mapping of code->text to translate error codes to readable text.
type ErrorSummary map[int]string

// ErrorFields defines extra fields for particular detailed error.
type ErrorFields map[string]any

// Error satisfies error interface for SynologyError type.
func (se ApiError) Error() string {
	buf := strings.Builder{}
	buf.WriteString(fmt.Sprintf("[%d] %s", se.Code, se.Summary))
	if len(se.Errors) > 0 {
		buf.WriteString("\n\tDetails:")
	}

	for _, e := range se.Errors {
		detailedFields := []string{}
		buf.WriteString(fmt.Sprintf("\n\t\t[%d] %s", e.Code, e.Summary))
		if len(e.Details) > 0 {
			for k, v := range e.Details {
				detailedFields = append(detailedFields, k+": "+fmt.Sprintf("%v", v))
			}
			buf.WriteString(": [" + strings.Join(detailedFields, ",") + "]")
		}
	}

	return buf.String()
}

// DescribeError translates error code to human-readable summary text.
// It accepts error code and number of summary maps to look in.
// First summary with this code wins.
func DescribeError(code int, summaries ...ErrorSummary) string {
	for _, summaryMap := range summaries {
		if summary, ok := summaryMap[code]; ok {
			return summary
		}
	}

	return "Unknown error code"
}

// UnmarshalJSON fullfills Unmarshaler interface for JSON objects.
func (ei *ErrorItem) UnmarshalJSON(b []byte) error {
	fields := map[string]any{}
	err := json.Unmarshal(b, &fields)
	if err != nil {
		return err
	}

	var code int

	if c, ok := fields["code"]; ok {
		if cc, ok := c.(float64); ok {
			code = int(cc)
		}
	}

	result := ErrorItem{
		Code: code,
	}
	if len(fields) > 0 {
		result.Details = ErrorFields{}
		for k, v := range fields {
			result.Details[k] = v
		}
	}
	*ei = result

	return nil
}
