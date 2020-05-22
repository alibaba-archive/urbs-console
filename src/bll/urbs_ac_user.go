package bll

import (
	"context"

	"github.com/teambition/urbs-console/src/dao"
	"github.com/teambition/urbs-console/src/schema"
	"github.com/teambition/urbs-console/src/tpl"
)

// UrbsAcUser ...
type UrbsAcUser struct {
	daos *dao.Daos
}

// List 返回用户列表
func (a *UrbsAcUser) List(ctx context.Context, args *tpl.ConsolePagination) (*tpl.UrbsAcUserListRes, error) {
	items, err := a.daos.UrbsAcUser.List(ctx, args)
	if err != nil {
		return nil, err
	}
	total, err := a.daos.UrbsAcUser.Count(ctx)
	if err != nil {
		return nil, err
	}
	res := &tpl.UrbsAcUserListRes{Result: items}
	res.TotalSize = total
	if len(res.Result) > args.PageSize {
		res.NextPageToken = args.GetNextPageToken()
		res.Result = items[:args.PageSize]
	}
	return res, nil
}

// Search ...
func (a *UrbsAcUser) Search(ctx context.Context, key string) (*tpl.UrbsAcUserListRes, error) {
	items, err := a.daos.UrbsAcUser.Search(ctx, key)
	if err != nil {
		return nil, err
	}
	res := &tpl.UrbsAcUserListRes{Result: items}
	return res, nil
}

// Add ...
func (a *UrbsAcUser) Add(ctx context.Context, body *tpl.UrbsAcUsersBody) error {
	users := make([]*schema.UrbsAcUser, len(body.Users))
	for i, user := range body.Users {
		users[i] = &schema.UrbsAcUser{
			UID:  user.Uid,
			Name: user.Name,
		}
	}
	return a.daos.UrbsAcUser.BatchAdd(ctx, users)
}

// DeleteByUID ...
func (a *UrbsAcUser) DeleteByUID(ctx context.Context, uid string) error {
	return a.daos.UrbsAcUser.DeleteByUID(ctx, uid)
}

// UpdateByUID ...
func (a *UrbsAcUser) UpdateByUID(ctx context.Context, name, uid string) error {
	return a.daos.UrbsAcUser.UpdateByUID(ctx, name, uid)
}
