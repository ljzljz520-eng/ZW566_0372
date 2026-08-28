package report

import (
	"fmt"
	"repairdesk/model"
	"strings"
)

func Render(s Summary) string {
	return fmt.Sprintf("total=%d open=%d working=%d closed=%d archived=%d rate=%.2f", s.Total, s.Open, s.Working, s.Closed, s.Archived, CompletionRate(s))
}
func GroupByAssignee(rs []model.Record) map[string][]model.Record {
	out := map[string][]model.Record{}
	for _, r := range rs {
		out[r.Assignee] = append(out[r.Assignee], r)
	}
	return out
}
func Keywords(rs []model.Record) []string {
	seen := map[string]bool{}
	for _, r := range rs {
		for _, w := range strings.Fields(strings.ToLower(r.Title)) {
			seen[w] = true
		}
	}
	out := []string{}
	for w := range seen {
		out = append(out, w)
	}
	return out
}
func Overdue(rs []model.Record, cutoff int64) []model.Record {
	out := []model.Record{}
	for _, r := range rs {
		if r.UpdatedAt.Unix() < cutoff && !model.IsTerminal(r.Status) {
			out = append(out, r)
		}
	}
	return out
}
