package model

import "strings"

type Query struct {
	Term, Status, Priority, Assignee string
	Limit                            int
}

func (q Query) Match(r Record) bool {
	if q.Status != "" && r.Status != q.Status {
		return false
	}
	if q.Priority != "" && r.Priority != q.Priority {
		return false
	}
	if q.Assignee != "" && r.Assignee != q.Assignee {
		return false
	}
	if q.Term != "" {
		t := strings.ToLower(q.Term)
		if !strings.Contains(strings.ToLower(r.Title), t) && !strings.Contains(strings.ToLower(r.Description), t) {
			return false
		}
	}
	return true
}
func (q Query) Normalize() Query {
	q.Term = strings.TrimSpace(q.Term)
	q.Status = strings.TrimSpace(strings.ToLower(q.Status))
	q.Priority = NormalizePriority(q.Priority)
	if q.Limit < 0 {
		q.Limit = 0
	}
	return q
}
func ApplyQuery(rs []Record, q Query) []Record {
	q = q.Normalize()
	out := []Record{}
	for _, r := range rs {
		if q.Match(r) {
			out = append(out, r)
			if q.Limit > 0 && len(out) >= q.Limit {
				break
			}
		}
	}
	return out
}
func ValidStatus(s string) bool {
	for _, x := range Statuses() {
		if x == s {
			return true
		}
	}
	return false
}
func ValidPriority(s string) bool { return s == "low" || s == "normal" || s == "high" }
func MergeTags(a, b []string) []string {
	out := append([]string{}, a...)
	for _, x := range b {
		found := false
		for _, y := range out {
			if x == y {
				found = true
				break
			}
		}
		if !found {
			out = append(out, x)
		}
	}
	return out
}
