package core

import (
	"strings"
	"testing"

	"github.com/synology-community/go-synology/pkg/query"
)

// PLAT-522: TaskRequest.Enable must reach the wire as an explicit
// enable=false, not be dropped, so a plan that disables an already-enabled
// task can actually take effect. `omitempty` on a bool field causes the
// go-querystring encoder to silently drop `false` values entirely, which
// previously made disabling a task impossible.
func TestTaskRequest_EncodeValues_EnableFalse(t *testing.T) {
	req := TaskRequest{
		Name:   "test",
		Enable: false,
	}

	v, err := query.Values(req)
	if err != nil {
		t.Fatalf("query.Values() error = %v", err)
	}

	if got := v.Get("enable"); got != "false" {
		t.Errorf("TaskRequest.EncodeValues() enable = %q, want %q", got, "false")
	}

	encoded := v.Encode()
	if !strings.Contains(encoded, "enable=false") {
		t.Errorf("TaskRequest.EncodeValues() encoded = %q, want it to contain %q", encoded, "enable=false")
	}
}

// TestTaskRequest_EncodeValues_EnableTrue guards the other direction: a
// true value must still be sent (this passed before the fix too, but is
// worth pinning so a future regression can't flip both cases silently).
func TestTaskRequest_EncodeValues_EnableTrue(t *testing.T) {
	req := TaskRequest{
		Name:   "test",
		Enable: true,
	}

	v, err := query.Values(req)
	if err != nil {
		t.Fatalf("query.Values() error = %v", err)
	}

	encoded := v.Encode()
	if !strings.Contains(encoded, "enable=true") {
		t.Errorf("TaskRequest.EncodeValues() encoded = %q, want it to contain %q", encoded, "enable=true")
	}
}
