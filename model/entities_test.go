package model

import "testing"

func TestRecordValidation(t *testing.T) {
	r := Record{ID: "1", Title: "fix", Status: StatusOpen}
	if r.Valid() != nil {
		t.Fatal("valid")
	}
}
