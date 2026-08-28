package repairdesk

import (
	"net/http/httptest"
	"repairdesk/api"
	"repairdesk/service"
	"repairdesk/storage"
	"testing"
)

func TestHTTPWorkflow(t *testing.T) {
	s, _ := storage.Open(t.TempDir() + "/i.db")
	defer s.Close()
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/health", nil)
	(&api.Server{Desk: service.New(s)}).Handler().ServeHTTP(rr, req)
	if rr.Code != 200 {
		t.Fatal(rr.Code)
	}
}
