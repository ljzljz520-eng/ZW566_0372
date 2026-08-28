package report

import (
	"repairdesk/model"
	"testing"
)

func TestSummary(t *testing.T) {
	s := Build([]model.Record{{Status: model.StatusClosed, Priority: "high"}})
	if CompletionRate(s) != 1 {
		t.Fatal(s)
	}
}
