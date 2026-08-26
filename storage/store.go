package storage

import (
	"encoding/json"
	"go.etcd.io/bbolt"
	"os"
	"repairdesk/model"
	"sync"
)

var buckets = [][]byte{[]byte("records"), []byte("profiles"), []byte("events"), []byte("audits")}

type Store struct {
	db *bbolt.DB
	mu sync.RWMutex
}

func Open(path string) (*Store, error) {
	db, e := bbolt.Open(path, 0600, nil)
	if e != nil {
		return nil, e
	}
	s := &Store{db: db}
	e = db.Update(func(tx *bbolt.Tx) error {
		for _, b := range buckets {
			if _, x := tx.CreateBucketIfNotExists(b); x != nil {
				return x
			}
		}
		return nil
	})
	if e != nil {
		db.Close()
		return nil, e
	}
	return s, nil
}
func (s *Store) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.db == nil {
		return nil
	}
	e := s.db.Close()
	s.db = nil
	return e
}
func enc(v any) ([]byte, error) { return json.Marshal(v) }
func put(tx *bbolt.Tx, b, k []byte, v any) error {
	x, e := enc(v)
	if e != nil {
		return e
	}
	return tx.Bucket(b).Put(k, x)
}
func get(tx *bbolt.Tx, b, k []byte, v any) error {
	x := tx.Bucket(b).Get(k)
	if x == nil {
		return model.ErrNotFound
	}
	return json.Unmarshal(x, v)
}
func (s *Store) SaveRecord(r model.Record) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return os.ErrClosed
	}
	return s.db.Update(func(tx *bbolt.Tx) error { return put(tx, buckets[0], []byte(r.ID), r) })
}
func (s *Store) Record(id string) (model.Record, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var r model.Record
	if s.db == nil {
		return r, os.ErrClosed
	}
	e := s.db.View(func(tx *bbolt.Tx) error { return get(tx, buckets[0], []byte(id), &r) })
	return r, e
}
func (s *Store) SaveProfile(p model.Profile) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.db.Update(func(tx *bbolt.Tx) error { return put(tx, buckets[1], []byte(p.ID), p) })
}
func (s *Store) Profile(id string) (model.Profile, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var p model.Profile
	e := s.db.View(func(tx *bbolt.Tx) error { return get(tx, buckets[1], []byte(id), &p) })
	return p, e
}
func (s *Store) SaveEvent(v model.Event) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.db.Update(func(tx *bbolt.Tx) error { return put(tx, buckets[2], []byte(v.ID), v) })
}
func (s *Store) SaveAudit(v model.Audit) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.db.Update(func(tx *bbolt.Tx) error { return put(tx, buckets[3], []byte(v.ID), v) })
}
