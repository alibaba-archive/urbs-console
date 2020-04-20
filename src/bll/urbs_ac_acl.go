package bll

import (
	"context"

	"github.com/teambition/urbs-console/src/apperrs"
	"github.com/teambition/urbs-console/src/conf"
	"github.com/teambition/urbs-console/src/constant"
	"github.com/teambition/urbs-console/src/dao"
	"github.com/teambition/urbs-console/src/schema"
	"github.com/teambition/urbs-console/src/tpl"
	"github.com/teambition/urbs-console/src/util"
)

// UrbsAcAcl ...
type UrbsAcAcl struct {
	daos *dao.Daos
}

// AddByReq ...
func (a *UrbsAcAcl) AddByReq(ctx context.Context, req *tpl.UrbsAcAclAddReq) error {
	object := req.Product + req.Label + req.Module + req.Setting
	return a.Add(ctx, req.Uid, object, req.Permission)
}

// AddDefaultPermission ...
func (a *UrbsAcAcl) AddDefaultPermission(ctx context.Context, subject string, object string) error {
	return a.Add(ctx, subject, object, constant.PermissionAll)
}

// Add ...
func (a *UrbsAcAcl) Add(ctx context.Context, subject string, object string, permission string) error {
	obj := &schema.UrbsAcAcl{
		Subject:    subject,
		Object:     object,
		Permission: permission,
	}
	return a.daos.UrbsAcAcl.Add(ctx, obj)
}

// Update ...
func (a *UrbsAcAcl) Update(ctx context.Context, subjects []string, object string) error {
	for _, subject := range subjects {
		err := a.AddDefaultPermission(ctx, subject, object)
		if err != nil {
			return err
		}
	}
	err := a.daos.UrbsAcAcl.DeleteNotIn(ctx, subjects, object)
	if err != nil {
		return err
	}
	return nil
}

// CheckViewer ...
func (a *UrbsAcAcl) CheckViewer(ctx context.Context) error {
	uid := util.GetUid(ctx)
	_, err := a.daos.UrbsAcUser.FindByUID(ctx, uid)
	if err != nil {
		return apperrs.ErrForbidden.WithMsg(uid)
	}
	return nil
}

// CheckAdmin ...
func (a *UrbsAcAcl) CheckAdmin(ctx context.Context, object string) error {
	return a.Check(ctx, object, constant.PermissionAll)
}

// CheckSuperAdmin ...
func (a *UrbsAcAcl) CheckSuperAdmin(ctx context.Context) error {
	uid := util.GetUid(ctx)
	res := util.StringInSlice(uid, conf.Config.SuperAdmins)
	if !res {
		return apperrs.ErrForbidden.WithMsg(uid)
	}
	return nil
}

// Check ...
func (a *UrbsAcAcl) Check(ctx context.Context, object string, permission string) error {
	uid := util.GetUid(ctx)
	_, err := a.daos.UrbsAcAcl.FindOne(ctx, uid, object, permission)
	if err != nil {
		return apperrs.ErrForbidden.WithMsg(uid)
	}
	return nil
}

// FindUsersByObjects ...
func (a *UrbsAcAcl) FindUsersByObjects(ctx context.Context, objects []string) (map[string][]*tpl.User, error) {
	acls, err := a.daos.UrbsAcAcl.FindByObjects(ctx, objects)
	if err != nil {
		return nil, err
	}
	subjects := map[string][]*tpl.User{}
	for _, acl := range acls {
		vals, ok := subjects[acl.Object]
		if ok {
			vals = append(vals, &tpl.User{
				Name: acl.Name, Uid: acl.Subject,
			})
		} else {
			subjects[acl.Subject] = []*tpl.User{{
				Name: acl.Name, Uid: acl.Subject,
			}}
		}
	}
	return subjects, nil
}

// FindUsersByObject ...
func (a *UrbsAcAcl) FindUsersByObject(ctx context.Context, object string) ([]*tpl.User, error) {
	objects, err := a.FindUsersByObjects(ctx, []string{object})
	if err != nil {
		return nil, err
	}
	return objects[object], nil
}

// FindOne ...
func (a *UrbsAcAcl) FindOne(ctx context.Context, subject, object, permission string) (*schema.UrbsAcAcl, error) {
	return a.daos.UrbsAcAcl.FindOne(ctx, subject, object, permission)
}

// Remove ...
func (a *UrbsAcAcl) Remove(ctx context.Context, subject, object, permission string) error {
	return a.daos.UrbsAcAcl.Delete(ctx, subject, object, permission)
}
