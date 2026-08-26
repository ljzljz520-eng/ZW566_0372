package service

import (
	"context"
	"errors"
	"repairdesk/model"
)

func ValidateRecord(r model.Record) error {
	if e := r.Valid(); e != nil {
		return e
	}
	if !model.ValidStatus(r.Status) {
		return model.ErrInvalidStatus
	}
	if !model.ValidPriority(r.Priority) {
		return errors.New("invalid priority")
	}
	return nil
}
func (d *Desk) Validate(ctx context.Context, id string) error {
	r, e := d.Get(ctx, id)
	if e != nil {
		return e
	}
	return ValidateRecord(r)
}
func (d *Desk) Transition(ctx context.Context, id, to string) error {
	if !model.ValidStatus(to) {
		return model.ErrInvalidStatus
	}
	r, e := d.Get(ctx, id)
	if e != nil {
		return e
	}
	if !model.CanTransition(r.Status, to) {
		return model.ErrConflict
	}
	return d.move(ctx, id, to, "")
}
func (d *Desk) AssignMany(ctx context.Context, ids []string, profile string) error {
	if profile == "" {
		return model.ErrInvalidProfile
	}
	for _, id := range ids {
		if e := d.Assign(ctx, id, profile); e != nil {
			return e
		}
	}
	return nil
}
func (d *Desk) RequireOpen(ctx context.Context, id string) error {
	r, e := d.Get(ctx, id)
	if e != nil {
		return e
	}
	if r.Status != model.StatusOpen {
		return model.ErrConflict
	}
	return nil
}
