package storage

import (
	"repairdesk/model"
	"time"
)

func (s *Store) UpdateRecord(id string, fn func(*model.Record) error) error {
	r, e := s.Record(id)
	if e != nil {
		return e
	}
	if e = fn(&r); e != nil {
		return e
	}
	r.UpdatedAt = time.Now()
	return s.SaveRecord(r)
}
func (s *Store) MarkTag(id, tag string) error {
	return s.UpdateRecord(id, func(r *model.Record) error { model.AddTag(r, tag); return nil })
}
func (s *Store) SetAssignee(id, assignee string) error {
	return s.UpdateRecord(id, func(r *model.Record) error {
		if assignee == "" {
			return model.ErrInvalidProfile
		}
		r.Assignee = assignee
		return nil
	})
}
func (s *Store) FindByStatus(status string) ([]model.Record, error) {
	rs, e := s.Records()
	if e != nil {
		return nil, e
	}
	return model.ApplyQuery(rs, model.Query{Status: status}), nil
}
func (s *Store) FindByAssignee(a string) ([]model.Record, error) {
	rs, e := s.Records()
	if e != nil {
		return nil, e
	}
	return model.ApplyQuery(rs, model.Query{Assignee: a}), nil
}
func (s *Store) PurgeArchived(before time.Time) (int, error) {
	rs, e := s.Records()
	if e != nil {
		return 0, e
	}
	n := 0
	for _, r := range rs {
		if r.Status == model.StatusArchived && r.UpdatedAt.Before(before) {
			if e = s.DeleteRecord(r.ID); e != nil {
				return n, e
			}
			n++
		}
	}
	return n, nil
}
