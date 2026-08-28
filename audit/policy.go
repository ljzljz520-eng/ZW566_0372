package audit

import (
	"errors"
	"repairdesk/model"
	"strings"
)

func AllowedAction(action string) bool {
	switch strings.ToLower(action) {
	case "create", "assign", "start", "close", "archive", "reopen":
		return true
	}
	return false
}
func Check(a model.Audit) error {
	if e := a.Valid(); e != nil {
		return e
	}
	if !AllowedAction(a.Action) {
		return errors.New("action not allowed")
	}
	if a.Actor == "" {
		return errors.New("actor required")
	}
	return nil
}
func Summarize(as []model.Audit) map[string]int {
	out := map[string]int{}
	for _, a := range as {
		out[a.Action]++
	}
	return out
}
func Last(as []model.Audit) model.Audit {
	var out model.Audit
	for _, a := range as {
		if a.At.After(out.At) {
			out = a
		}
	}
	return out
}
