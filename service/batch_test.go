package service

import (
	"context"
	"repairdesk/storage"
	"testing"
)

func TestBusinessChain11(t *testing.T) {
	s, _ := storage.Open(t.TempDir() + "/b.db")
	defer s.Close()
	d := New(s)
	ids := []string{}
	for i := 0; i < 10; i++ {
		r, _ := d.Register(context.Background(), "job", "desc", "normal")
		d.Assign(context.Background(), r.ID, "p")
		d.Start(context.Background(), r.ID)
		ids = append(ids, r.ID)
	}
	if e := d.BatchClose(context.Background(), ids); e != nil {
		t.Fatal(e)
	}
	if d.ActiveCount() != 0 {
		t.Fatalf("handles remain: %d", d.ActiveCount())
	}
}

func TestBatchReleasesEachHandle(t *testing.T) {
	s, _ := storage.Open(t.TempDir() + "/c.db")
	defer s.Close()
	d := New(s)
	ids := []string{}
	for i := 0; i < 10; i++ {
		r, _ := d.Register(context.Background(), "job", "desc", "normal")
		d.Assign(context.Background(), r.ID, "p")
		d.Start(context.Background(), r.ID)
		ids = append(ids, r.ID)
	}
	var observed []int
	if e := d.BatchCloseObserved(context.Background(), ids, func(_ int) { observed = append(observed, d.ActiveCount()) }); e != nil {
		t.Fatal(e)
	}
	for _, n := range observed {
		if n != 0 {
			t.Fatalf("resource remained active: %d", n)
		}
	}
}
