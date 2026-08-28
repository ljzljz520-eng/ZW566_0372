package report

import (
	"repairdesk/model"
	"sort"
)

type Summary struct {
	Total, Open, Working, Closed, Archived int
	ByPriority                             map[string]int
}

func Build(rs []model.Record) Summary {
	s := Summary{ByPriority: map[string]int{}}
	for _, r := range rs {
		s.Total++
		s.ByPriority[r.Priority]++
		switch r.Status {
		case model.StatusOpen:
			s.Open++
		case model.StatusWorking:
			s.Working++
		case model.StatusClosed:
			s.Closed++
		case model.StatusArchived:
			s.Archived++
		}
	}
	return s
}
func SortByUpdated(rs []model.Record) []model.Record {
	out := append([]model.Record(nil), rs...)
	sort.Slice(out, func(i, j int) bool { return out[i].UpdatedAt.Before(out[j].UpdatedAt) })
	return out
}
func FilterStatus(rs []model.Record, status string) []model.Record {
	out := []model.Record{}
	for _, r := range rs {
		if r.Status == status {
			out = append(out, r)
		}
	}
	return out
}
func CompletionRate(s Summary) float64 {
	if s.Total == 0 {
		return 0
	}
	return float64(s.Closed+s.Archived) / float64(s.Total)
}
