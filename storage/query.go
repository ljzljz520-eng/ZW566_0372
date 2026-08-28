package storage

import (
	"encoding/json"
	"go.etcd.io/bbolt"
	"repairdesk/model"
)

func list[T any](s *Store, b []byte) ([]T, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := []T{}
	e := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket(b).ForEach(func(_, v []byte) error {
			var x T
			if e := json.Unmarshal(v, &x); e != nil {
				return e
			}
			out = append(out, x)
			return nil
		})
	})
	return out, e
}
func (s *Store) Records() ([]model.Record, error)   { return list[model.Record](s, buckets[0]) }
func (s *Store) Profiles() ([]model.Profile, error) { return list[model.Profile](s, buckets[1]) }
func (s *Store) Events() ([]model.Event, error)     { return list[model.Event](s, buckets[2]) }
func (s *Store) Audits() ([]model.Audit, error)     { return list[model.Audit](s, buckets[3]) }
func (s *Store) DeleteRecord(id string) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket(buckets[0]).Delete([]byte(id)) })
}
