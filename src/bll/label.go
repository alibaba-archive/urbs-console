package bll

import (
	"context"

	"github.com/teambition/urbs-console/src/dto/urbssetting"
	"github.com/teambition/urbs-console/src/service"
	"github.com/teambition/urbs-console/src/tpl"
)

// Label ...
type Label struct {
	services *service.Services
}

// Create ...
func (a *Label) Create(ctx context.Context, productName string, args *tpl.LabelBody) (*urbssetting.LabelInfoRes, error) {
	return a.services.UrbsSetting.LabelCreate(ctx, productName, args)
}

// List 返回产品下的标签列表
func (a *Label) List(ctx context.Context, req *tpl.ProductPaginationURL) (*urbssetting.LabelsInfoRes, error) {
	return a.services.UrbsSetting.LabelList(ctx, req)
}

// Update ...
func (a *Label) Update(ctx context.Context, productName, labelName string, body *tpl.LabelUpdateBody) (*urbssetting.LabelInfoRes, error) {
	return a.services.UrbsSetting.LabelUpdate(ctx, productName, labelName, body)
}

// Offline 下线标签
func (a *Label) Offline(ctx context.Context, productName, labelName string) (*tpl.BoolRes, error) {
	return a.services.UrbsSetting.LabelOffline(ctx, productName, labelName)
}

// Assign 把标签批量分配给用户或群组
func (a *Label) Assign(ctx context.Context, productName, labelName string, body *tpl.UsersGroupsBody) (*tpl.BoolRes, error) {
	return a.services.UrbsSetting.LabelAssign(ctx, productName, labelName, body)
}

// Delete 物理删除标签
func (a *Label) Delete(ctx context.Context, productName, labelName string) (*tpl.BoolRes, error) {
	return a.services.UrbsSetting.LabelDelete(ctx, productName, labelName)
}
