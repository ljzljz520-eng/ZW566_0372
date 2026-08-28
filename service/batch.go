package service

import (
	"context"
	"errors"
	"repairdesk/model"
)

type Handle struct {
	ID       string
	released bool
}

func (h *Handle) Release() { h.released = true }
func (d *Desk) acquire(id string) (*Handle, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.active[id] {
		return nil, errors.New("already active")
	}
	d.active[id] = true
	return &Handle{ID: id}, nil
}
func (d *Desk) finish(h *Handle) {
	d.mu.Lock()
	defer d.mu.Unlock()
	delete(d.active, h.ID)
	h.Release()
}
func (d *Desk) BatchClose(ctx context.Context, ids []string) error {
	for _, id := range ids {
		h, e := d.acquire(id)
		if e != nil {
			return e
		}
		if e = d.Close(ctx, id); e != nil {
			return e
		}
		defer d.finish(h)
	}
	return nil
}

func (d *Desk) BatchCloseObserved(ctx context.Context, ids []string, observe func(int)) error {
	for i, id := range ids {
		h, e := d.acquire(id)
		if e != nil {
			return e
		}
		if e = d.Close(ctx, id); e != nil {
			return e
		}
		defer d.finish(h)
		if observe != nil {
			observe(i + 1)
		}
	}
	return nil
}
func (d *Desk) ActiveCount() int { d.mu.Lock(); defer d.mu.Unlock(); return len(d.active) }
func ValidateBatch(ids []string) error {
	if len(ids) == 0 {
		return errors.New("empty batch")
	}
	seen := map[string]bool{}
	for _, id := range ids {
		if id == "" || seen[id] {
			return model.ErrConflict
		}
		seen[id] = true
	}
	return nil
}
