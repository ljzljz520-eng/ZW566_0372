package workflow

import (
	"context"
	"repairdesk/model"
	"repairdesk/service"
)

type Chain struct{ desk *service.Desk }

func New(d *service.Desk) *Chain { return &Chain{desk: d} }
func (c *Chain) Intake(ctx context.Context, title, desc string) (model.Record, error) {
	return c.desk.Register(ctx, title, desc, "normal")
}
func (c *Chain) Process(ctx context.Context, id, assignee string) error {
	if e := c.desk.Assign(ctx, id, assignee); e != nil {
		return e
	}
	return c.desk.Start(ctx, id)
}
func (c *Chain) Complete(ctx context.Context, id string) error {
	if e := c.desk.Close(ctx, id); e != nil {
		return e
	}
	return c.desk.Archive(ctx, id)
}
func (c *Chain) Full(ctx context.Context, title, desc, assignee string) (model.Record, error) {
	r, e := c.Intake(ctx, title, desc)
	if e != nil {
		return r, e
	}
	if e = c.Process(ctx, r.ID, assignee); e != nil {
		return r, e
	}
	if e = c.Complete(ctx, r.ID); e != nil {
		return r, e
	}
	return c.desk.Get(ctx, r.ID)
}
