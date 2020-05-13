package bll

import (
	"context"

	"github.com/teambition/urbs-console/src/dto"
	"github.com/teambition/urbs-console/src/logger"
	"github.com/teambition/urbs-console/src/service"
	"github.com/teambition/urbs-console/src/tpl"
)

// Label ...
type Label struct {
	services *service.Services
}

// ListGroups ...
func (a *Label) ListGroups(ctx context.Context, args *tpl.ProductLabelURL) (*tpl.LabelGroupsInfoRes, error) {
	return a.services.UrbsSetting.LabelListGroups(ctx, args)
}

// ListUsers ...
func (a *Label) ListUsers(ctx context.Context, args *tpl.ProductLabelURL) (*tpl.LabelUsersInfoRes, error) {
	return a.services.UrbsSetting.LabelListUsers(ctx, args)
}

// Create ...
func (a *Label) Create(ctx context.Context, product string, args *tpl.LabelBody) (*tpl.LabelInfoRes, error) {
	aclObject := product + args.Name
	err := blls.UrbsAcAcl.AddDefaultPermission(ctx, args.Uids, aclObject)
	if err != nil {
		return nil, err
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
	err := blls.UrbsAcAcl.Update(ctx, body.Uids, product+label)
	if err != nil {
		return nil, err
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
func (a *Label) Assign(ctx context.Context, args *tpl.ProductLabelURL, body *tpl.UsersGroupsBody) (*tpl.LabelReleaseInfoRes, error) {
	object := args.Product + args.Label
	logContent := &dto.OperationLogContent{
		Users:  body.Users,
		Groups: body.Groups,
		Desc:   body.Desc,
		Value:  body.Value,
	}
	err := blls.OperationLog.Add(ctx, object, actionCreate, logContent)
	if err != nil {
		return nil, err
	}
	return a.services.UrbsSetting.LabelAssign(ctx, args.Product, args.Label, body)
}

// Delete 物理删除标签
func (a *Label) Delete(ctx context.Context, product, label string) (*tpl.BoolRes, error) {
	err := daos.UrbsAcAcl.DeleteByObject(ctx, product+label)
	if err != nil {
		logger.Err(ctx, err.Error())
	}
	return a.services.UrbsSetting.LabelDelete(ctx, product, label)
}

// Recall 批量撤销对用户或群组设置的产品灰度标签
func (a *Label) Recall(ctx context.Context, args *tpl.ProductLabelURL, body *tpl.RecallBody) (*tpl.BoolRes, error) {
	object := args.Product + args.Label
	err := daos.OperationLog.DeleteByObject(ctx, object)
	if err != nil {
		return nil, err
	}
	return a.services.UrbsSetting.LabelRecall(ctx, args, body)
}

// ListRules ...
func (a *Label) ListRules(ctx context.Context, args *tpl.ProductLabelURL) (*tpl.LabelRulesInfoRes, error) {
	return a.services.UrbsSetting.LabelListRule(ctx, args)
}

// CreateRule ...
func (a *Label) CreateRule(ctx context.Context, args *tpl.ProductLabelURL, body *tpl.LabelRuleBody) (*tpl.LabelRuleInfoRes, error) {
	object := args.Product + args.Label
	logContent := &dto.OperationLogContent{
		Desc:    body.Desc,
		Percent: body.Rule.Value,
	}
	err := blls.OperationLog.Add(ctx, object, actionCreate, logContent)
	if err != nil {
		return nil, err
	}
	return a.services.UrbsSetting.LabelCreateRule(ctx, args, body)
}

// UpdateRule ...
func (a *Label) UpdateRule(ctx context.Context, args *tpl.ProductLabelHIDURL, body *tpl.LabelRuleBody) (*tpl.LabelRuleInfoRes, error) {
	object := args.Product + args.Label
	logContent := &dto.OperationLogContent{
		Desc:    body.Desc,
		Percent: body.Rule.Value,
	}
	err := blls.OperationLog.Add(ctx, object, actionUpdate, logContent)
	if err != nil {
		return nil, err
	}
	return a.services.UrbsSetting.LabelUpdateRule(ctx, args, body)
}

// DeleteRule ...
func (a *Label) DeleteRule(ctx context.Context, args *tpl.ProductLabelHIDURL) (*tpl.BoolRes, error) {
	return a.services.UrbsSetting.LabelDeleteRule(ctx, args)
}
