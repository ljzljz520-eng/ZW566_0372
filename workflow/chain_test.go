package workflow

import (
	"context"
	"repairdesk/service"
	"repairdesk/storage"
	"testing"
)

func TestWorkflowOne(t *testing.T) {
	s, _ := storage.Open(t.TempDir() + "/w.db")
	defer s.Close()
	r, e := New(service.New(s)).Intake(context.Background(), "new", "desc")
	if e != nil || r.Status != "open" {
		t.Fatal(e)
	}
}
func TestWorkflowTwo(t *testing.T) {
	s, _ := storage.Open(t.TempDir() + "/w2.db")
	defer s.Close()
	c := New(service.New(s))
	r, _ := c.Intake(context.Background(), "new", "desc")
	if e := c.Process(context.Background(), r.ID, "p"); e != nil {
		t.Fatal(e)
	}
}
func TestWorkflowThree(t *testing.T) {
	s, _ := storage.Open(t.TempDir() + "/w3.db")
	defer s.Close()
	c := New(service.New(s))
	r, e := c.Full(context.Background(), "new", "desc", "p")
	if e != nil || r.Status != "archived" {
		t.Fatal(e, r.Status)
	}
}
