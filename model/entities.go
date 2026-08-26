package model

import "time"

type Record struct {
	ID, Title, Description, Status, Priority, Assignee string
	CreatedAt, UpdatedAt                               time.Time
	Tags                                               []string
}
type Profile struct {
	ID, Name, Team, Phone string
	Active                bool
	Skills                []string
}
type Event struct {
	ID, RecordID, Kind, Actor, Detail string
	At                                time.Time
}
type Audit struct {
	ID, RecordID, Action, Actor, Result string
	At                                  time.Time
}

const (
	StatusOpen     = "open"
	StatusAssigned = "assigned"
	StatusWorking  = "working"
	StatusClosed   = "closed"
	StatusArchived = "archived"
)

func (r Record) Valid() error {
	if r.ID == "" || r.Title == "" {
		return ErrInvalidRecord
	}
	if r.Status == "" {
		return ErrInvalidStatus
	}
	return nil
}
func (p Profile) Valid() error {
	if p.ID == "" || p.Name == "" {
		return ErrInvalidProfile
	}
	return nil
}
func (e Event) Valid() error {
	if e.ID == "" || e.RecordID == "" || e.Kind == "" {
		return ErrInvalidEvent
	}
	return nil
}
func (a Audit) Valid() error {
	if a.ID == "" || a.RecordID == "" || a.Action == "" {
		return ErrInvalidAudit
	}
	return nil
}
