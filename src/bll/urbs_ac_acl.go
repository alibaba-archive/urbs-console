package bll

import (
	"context"

	"github.com/teambition/urbs-console/src/dao"
	"github.com/teambition/urbs-console/src/schema"
)

// UrbsAcAcl ...
type UrbsAcAcl struct {
	daos *dao.Daos
}

// Add ...
func (a *UrbsAcAcl) Add(ctx context.Context, obj *schema.UrbsAcAcl) error {
	return a.daos.UrbsAcAcl.Add(ctx, obj)
}

// FindOne ...
func (a *UrbsAcAcl) FindOne(ctx context.Context, subject, object, permission string) (*schema.UrbsAcAcl, error) {
	return a.daos.UrbsAcAcl.FindOne(ctx, subject, object, permission)
}

// Remove ...
func (a *UrbsAcAcl) Remove(ctx context.Context, subject, object, permission string) error {
	return a.daos.UrbsAcAcl.Remove(ctx, subject, object, permission)
}
