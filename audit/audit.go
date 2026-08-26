package audit

import (
	"fmt"
	"repairdesk/model"
	"repairdesk/storage"
	"time"
)

type Logger struct{ store *storage.Store }

func New(s *storage.Store) *Logger { return &Logger{store: s} }
func (l *Logger) Record(recordID, action, actor, result string) error {
	a := model.Audit{ID: fmt.Sprintf("A-%d", time.Now().UnixNano()), RecordID: recordID, Action: action, Actor: actor, Result: result, At: time.Now()}
	if e := a.Valid(); e != nil {
		return e
	}
	return l.store.SaveAudit(a)
}
func (l *Logger) Event(recordID, kind, actor, detail string) error {
	e := model.Event{ID: fmt.Sprintf("E-%d", time.Now().UnixNano()), RecordID: recordID, Kind: kind, Actor: actor, Detail: detail, At: time.Now()}
	if x := e.Valid(); x != nil {
		return x
	}
	return l.store.SaveEvent(e)
}
func (l *Logger) ForRecord(id string) ([]model.Audit, error) {
	xs, e := l.store.Audits()
	if e != nil {
		return nil, e
	}
	out := []model.Audit{}
	for _, a := range xs {
		if a.RecordID == id {
			out = append(out, a)
		}
	}
	return out, nil
}
func (l *Logger) EventsFor(id string) ([]model.Event, error) {
	xs, e := l.store.Events()
	if e != nil {
		return nil, e
	}
	out := []model.Event{}
	for _, v := range xs {
		if v.RecordID == id {
			out = append(out, v)
		}
	}
	return out, nil
}
