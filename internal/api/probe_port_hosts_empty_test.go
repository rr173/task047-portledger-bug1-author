package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"task047-portledger/internal/registry"
)

func TestProbePortHostsEmptyReturnsArray(t *testing.T) {
	rr := httptest.NewRecorder()
	New(registry.New()).ServeHTTP(rr, httptest.NewRequest(http.MethodGet, "/ports/443/hosts", nil))
	if rr.Code != http.StatusOK {
		t.Fatalf("status=%d want 200; body=%s", rr.Code, rr.Body.String())
	}
	var body struct {
		Hosts []string `json:"hosts"`
	}
	if err := json.Unmarshal(rr.Body.Bytes(), &body); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if body.Hosts == nil {
		t.Fatal("hosts must be an empty JSON array, not null")
	}
	if len(body.Hosts) != 0 {
		t.Fatalf("hosts=%v want empty", body.Hosts)
	}
}
