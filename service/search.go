package service

import (
	"context"
	"repairdesk/model"
	"repairdesk/report"
	"strings"
)

func (d *Desk) Search(ctx context.Context, term string) ([]model.Record, error) {
	rs, e := d.List(ctx)
	if e != nil {
		return nil, e
	}
	term = strings.ToLower(strings.TrimSpace(term))
	out := []model.Record{}
	for _, r := range rs {
		if term == "" || strings.Contains(strings.ToLower(r.Title), term) || strings.Contains(strings.ToLower(r.Description), term) {
			out = append(out, r)
		}
	}
	return report.SortByUpdated(out), nil
}
func (d *Desk) ByPriority(ctx context.Context, p string) ([]model.Record, error) {
	rs, e := d.List(ctx)
	if e != nil {
		return nil, e
	}
	out := []model.Record{}
	for _, r := range rs {
		if r.Priority == p {
			out = append(out, r)
		}
	}
	return out, nil
}
func (d *Desk) Reopen(ctx context.Context, id string) error {
	if e := ctx.Err(); e != nil {
		return e
	}
	d.mu.Lock()
	defer d.mu.Unlock()
	r, e := d.store.Record(id)
	if e != nil {
		return e
	}
	if r.Status != model.StatusClosed {
		return model.ErrConflict
	}
	r.Status = model.StatusWorking
	return d.store.SaveRecord(r)
}
func (d *Desk) EnsureProfile(p model.Profile) error {
	if e := p.Valid(); e != nil {
		return e
	}
	return d.store.SaveProfile(p)
}
