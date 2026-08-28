package storage

import (
	"repairdesk/model"
	"testing"
)

func TestQueryRecords(t *testing.T) {
	s, e := Open(t.TempDir() + "/q.db")
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	s.SaveRecord(model.Record{ID: "1", Title: "a", Status: model.StatusOpen, Priority: "normal"})
	xs, e := s.FindByStatus(model.StatusOpen)
	if e != nil || len(xs) != 1 {
		t.Fatal(len(xs), e)
	}
}
