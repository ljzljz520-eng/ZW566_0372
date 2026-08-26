package service

import (
	"context"
	"repairdesk/storage"
	"testing"
)

func TestRegisterAndAssign(t *testing.T) {
	s, _ := storage.Open(t.TempDir() + "/s.db")
	defer s.Close()
	d := New(s)
	r, e := d.Register(context.Background(), "leak", "pump", "high")
	if e != nil {
		t.Fatal(e)
	}
	if e = d.Assign(context.Background(), r.ID, "p1"); e != nil {
		t.Fatal(e)
	}
}
