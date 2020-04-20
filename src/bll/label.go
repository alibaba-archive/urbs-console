package bll

import (
	"context"

	"github.com/teambition/urbs-console/src/logger"
	"github.com/teambition/urbs-console/src/service"
	"github.com/teambition/urbs-console/src/tpl"
)

// Label ...
type Label struct {
	services *service.Services
}

// GetGroups ...
func (a *Label) GetGroups(ctx context.Context, product, label string) (*tpl.LabelGroupsRes, error) {
	res := new(tpl.LabelGroupsRes)
	return res, nil
}

// GetUsers ...
func (a *Label) GetUsers(ctx context.Context, product, label string) (*tpl.LabelUsersRes, error) {
	res := new(tpl.LabelUsersRes)
	return res, nil
}

// Recall ...
func (a *Label) Recall(ctx context.Context, product, label string) error {
	_, err := daos.OperationLog.FindOneByObject(ctx, product+label)
	if err != nil {
		return nil
	}
	// call urbs-setting
	return nil
}

// Create ...
func (a *Label) Create(ctx context.Context, product string, args *tpl.LabelBody) (*tpl.LabelInfoRes, error) {
	aclObject := product + args.Name
	for _, uid := range args.Uids {
		err := blls.UrbsAcAcl.AddDefaultPermission(ctx, uid, aclObject)
		if err != nil {
			return nil, err
		}
	}
	res, err := a.services.UrbsSetting.LabelCreate(ctx, product, args)
	if err != nil {
		return nil, err
	}
	res.Result.Users, err = blls.UrbsAcAcl.FindUsersByObject(ctx, aclObject)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// List 返回产品下的标签列表
func (a *Label) List(ctx context.Context, args *tpl.ProductPaginationURL) (*tpl.LabelsInfoRes, error) {
	labels, err := a.services.UrbsSetting.LabelList(ctx, args)
	if err != nil {
		return nil, err
	}
	objects := make([]string, len(labels.Result))
	for i, label := range labels.Result {
		objects[i] = args.Product + label.Name
	}
	subjects, err := blls.UrbsAcAcl.FindUsersByObjects(ctx, objects)
	if err != nil {
		return nil, err
	}
	for _, label := range labels.Result {
		label.Users = subjects[args.Product+label.Name]
	}
	return labels, nil
}

// Update ...
func (a *Label) Update(ctx context.Context, product, label string, body *tpl.LabelUpdateBody) (*tpl.LabelInfoRes, error) {
	aclObject := product + label
	if len(body.Uids) > 0 {
		err := blls.UrbsAcAcl.Update(ctx, body.Uids, product+label)
		if err != nil {
			return nil, err
		}
	}
	res, err := a.services.UrbsSetting.LabelUpdate(ctx, product, label, body)
	if err != nil {
		return nil, err
	}
	res.Result.Users, err = blls.UrbsAcAcl.FindUsersByObject(ctx, aclObject)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Offline 下线标签
func (a *Label) Offline(ctx context.Context, product, label string) (*tpl.BoolRes, error) {
	return a.services.UrbsSetting.LabelOffline(ctx, product, label)
}

// Assign 把标签批量分配给用户或群组
func (a *Label) Assign(ctx context.Context, product, label string, body *tpl.UsersGroupsBody) (*tpl.BoolRes, error) {
	err := blls.OperationLog.Add(ctx, product+label, actionCreate, body)
	if err != nil {
		return nil, err
	}
	return a.services.UrbsSetting.LabelAssign(ctx, product, label, body)
}

// Delete 物理删除标签
func (a *Label) Delete(ctx context.Context, product, label string) (*tpl.BoolRes, error) {
	err := daos.UrbsAcAcl.DeleteByObject(ctx, product+label)
	if err != nil {
		logger.Err(ctx, err.Error())
	}
	return a.services.UrbsSetting.LabelDelete(ctx, product, label)
}
