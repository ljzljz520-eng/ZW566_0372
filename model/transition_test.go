package model

import "testing"

func TestTransitions(t *testing.T) {
	if !CanTransition(StatusOpen, StatusAssigned) || CanTransition(StatusOpen, StatusClosed) {
		t.Fatal("transition")
	}
}
