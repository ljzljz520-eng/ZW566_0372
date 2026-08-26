package report

import (
	"repairdesk/model"
	"sort"
)

type Metrics struct {
	ByStatus   map[string]int
	ByAssignee map[string]int
	Priorities []string
}

func Compute(rs []model.Record) Metrics {
	m := Metrics{ByStatus: map[string]int{}, ByAssignee: map[string]int{}}
	for _, r := range rs {
		m.ByStatus[r.Status]++
		m.ByAssignee[r.Assignee]++
	}
	for p := range map[string]bool{"low": true, "normal": true, "high": true} {
		m.Priorities = append(m.Priorities, p)
	}
	sort.Strings(m.Priorities)
	return m
}
func OpenCount(rs []model.Record) int {
	n := 0
	for _, r := range rs {
		if !model.IsTerminal(r.Status) {
			n++
		}
	}
	return n
}
func PriorityCounts(rs []model.Record) map[string]int {
	m := map[string]int{}
	for _, r := range rs {
		m[r.Priority]++
	}
	return m
}
func Latest(rs []model.Record) model.Record {
	if len(rs) == 0 {
		return model.Record{}
	}
	out := rs[0]
	for _, r := range rs[1:] {
		if r.UpdatedAt.After(out.UpdatedAt) {
			out = r
		}
	}
	return out
}
