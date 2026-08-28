package workflow

import (
	"context"
	"repairdesk/model"
	"repairdesk/service"
)

func (c *Chain) IntakeMany(ctx context.Context, titles []string) ([]model.Record, error) {
	out := make([]model.Record, 0, len(titles))
	for _, t := range titles {
		r, e := c.Intake(ctx, t, "bulk intake")
		if e != nil {
			return out, e
		}
		out = append(out, r)
	}
	return out, nil
}
func (c *Chain) ProcessMany(ctx context.Context, ids []string, assignee string) error {
	for _, id := range ids {
		if e := c.Process(ctx, id, assignee); e != nil {
			return e
		}
	}
	return nil
}
func (c *Chain) CompleteMany(ctx context.Context, ids []string) error {
	if e := service.ValidateBatch(ids); e != nil {
		return e
	}
	for _, id := range ids {
		if e := c.Complete(ctx, id); e != nil {
			return e
		}
	}
	return nil
}
func (c *Chain) Verify(ctx context.Context, id string) bool {
	r, e := c.desk.Get(ctx, id)
	return e == nil && r.Status == model.StatusArchived
}
