package service

import (
	"context"
	"fmt"
	"repairdesk/model"
	"repairdesk/storage"
	"sync"
	"time"
)

type Desk struct {
	store  *storage.Store
	mu     sync.Mutex
	active map[string]bool
}

func New(s *storage.Store) *Desk { return &Desk{store: s, active: map[string]bool{}} }
func (d *Desk) Register(ctx context.Context, title, desc, priority string) (model.Record, error) {
	if err := ctx.Err(); err != nil {
		return model.Record{}, err
	}
	r := model.Record{ID: fmt.Sprintf("R-%d", time.Now().UnixNano()), Title: title, Description: desc, Priority: priority, Status: model.StatusOpen, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	if e := r.Valid(); e != nil {
		return r, e
	}
	return r, d.store.SaveRecord(r)
}
func (d *Desk) Assign(ctx context.Context, id, profile string) error {
	return d.move(ctx, id, model.StatusAssigned, profile)
}
func (d *Desk) Start(ctx context.Context, id string) error {
	return d.move(ctx, id, model.StatusWorking, "")
}
func (d *Desk) Close(ctx context.Context, id string) error {
	return d.move(ctx, id, model.StatusClosed, "")
}
func (d *Desk) Archive(ctx context.Context, id string) error {
	return d.move(ctx, id, model.StatusArchived, "")
}
func (d *Desk) move(ctx context.Context, id, to, actor string) error {
	if e := ctx.Err(); e != nil {
		return e
	}
	d.mu.Lock()
	defer d.mu.Unlock()
	r, e := d.store.Record(id)
	if e != nil {
		return e
	}
	if !model.CanTransition(r.Status, to) {
		return model.ErrConflict
	}
	r.Status = to
	r.UpdatedAt = time.Now()
	if actor != "" {
		r.Assignee = actor
	}
	return d.store.SaveRecord(r)
}
func (d *Desk) Get(ctx context.Context, id string) (model.Record, error) {
	if e := ctx.Err(); e != nil {
		return model.Record{}, e
	}
	return d.store.Record(id)
}
func (d *Desk) List(ctx context.Context) ([]model.Record, error) {
	if e := ctx.Err(); e != nil {
		return nil, e
	}
	return d.store.Records()
}
