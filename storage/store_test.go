package storage

import (
	"os"
	"repairdesk/model"
	"testing"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	p := t.TempDir() + "/x.db"
	s, e := Open(p)
	if e != nil {
		t.Fatal(e)
	}
	r := model.Record{ID: "r1", Title: "pump", Status: model.StatusOpen, Priority: "normal"}
	if e = s.SaveRecord(r); e != nil {
		t.Fatal(e)
	}
	s.Close()
	s, e = Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	got, e := s.Record("r1")
	if e != nil || got.Title != "pump" {
		t.Fatal(got, e)
	}
	_ = os.Remove(p)
}
