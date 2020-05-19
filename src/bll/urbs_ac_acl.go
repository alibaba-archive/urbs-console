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
func (a *UrbsAcAcl) AddByReq(ctx context.Context, args *tpl.UrbsAcAclURL, req *tpl.UrbsAcAclAddBody) error {
	object := req.Product + req.Label + req.Module + req.Setting
	return a.Add(ctx, args.Uid, object, req.Permission)
}

// AddDefaultPermission ...
func (a *UrbsAcAcl) AddDefaultPermission(ctx context.Context, subjects []string, object string) error {
	for _, subject := range subjects {
		err := a.Add(ctx, subject, object, constant.PermissionAll)
		if err != nil {
			return err
		}
	}
	return nil
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
func (a *UrbsAcAcl) Update(ctx context.Context, subjects *[]string, object string) error {
	if subjects == nil || len(*subjects) == 0 {
		return nil
	}
	err := a.daos.UrbsAcAcl.UpdateSubjects(ctx, *subjects, object, constant.PermissionAll)
	return err
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
	res := util.StringSliceHas(conf.Config.SuperAdmins, uid)
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
			subjects[acl.Object] = append(vals, &tpl.User{
				Name: acl.Name, Uid: acl.Subject,
			})
		} else {
			subjects[acl.Object] = []*tpl.User{{
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

// Delete ...
func (a *UrbsAcAcl) Delete(ctx context.Context, args *tpl.UrbsAcAclURL, req *tpl.UrbsAcAclAddBody) error {
	object := req.Product + req.Label + req.Module + req.Setting
	return a.daos.UrbsAcAcl.Delete(ctx, args.Uid, object, req.Permission)
}
