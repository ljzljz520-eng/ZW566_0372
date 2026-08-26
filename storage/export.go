package storage

import (
	"encoding/json"
	"repairdesk/model"
	"sort"
)

type Snapshot struct {
	Records  []model.Record  `json:"records"`
	Profiles []model.Profile `json:"profiles"`
	Events   []model.Event   `json:"events"`
	Audits   []model.Audit   `json:"audits"`
}

func (s *Store) Snapshot() (Snapshot, error) {
	r, e := s.Records()
	if e != nil {
		return Snapshot{}, e
	}
	p, e := s.Profiles()
	if e != nil {
		return Snapshot{}, e
	}
	v, e := s.Events()
	if e != nil {
		return Snapshot{}, e
	}
	a, e := s.Audits()
	if e != nil {
		return Snapshot{}, e
	}
	sort.Slice(r, func(i, j int) bool { return r[i].ID < r[j].ID })
	return Snapshot{r, p, v, a}, nil
}
func (s *Store) ExportJSON() ([]byte, error) {
	x, e := s.Snapshot()
	if e != nil {
		return nil, e
	}
	return json.MarshalIndent(x, "", "  ")
}
func (s *Store) SaveEntities(r model.Record, p model.Profile, v model.Event, a model.Audit) error {
	if e := s.SaveRecord(r); e != nil {
		return e
	}
	if e := s.SaveProfile(p); e != nil {
		return e
	}
	if e := s.SaveEvent(v); e != nil {
		return e
	}
	return s.SaveAudit(a)
}
func (s *Store) CountAll() (int, error) {
	r, e := s.Records()
	if e != nil {
		return 0, e
	}
	p, e := s.Profiles()
	if e != nil {
		return 0, e
	}
	v, e := s.Events()
	if e != nil {
		return 0, e
	}
	a, e := s.Audits()
	return len(r) + len(p) + len(v) + len(a), e
}
