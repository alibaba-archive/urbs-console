package bll

import (
	"context"

	"github.com/teambition/gear"
	"github.com/teambition/urbs-console/src/dto"
	"github.com/teambition/urbs-console/src/logger"
	"github.com/teambition/urbs-console/src/service"
	"github.com/teambition/urbs-console/src/tpl"
	"github.com/teambition/urbs-console/src/util"
)

// Setting ...
type Setting struct {
	services *service.Services
	blls     *Blls
}

// ListByProduct ...
func (a *Setting) ListByProduct(ctx context.Context, args *tpl.ProductPaginationURL) (*tpl.SettingsInfoRes, error) {
	ress, err := a.services.UrbsSetting.SettingListByProduct(ctx, args)
	if err != nil {
		return nil, err
	}
	objects := make([]string, len(ress.Result))
	for i, setting := range ress.Result {
		objects[i] = args.Product + setting.Module + setting.Name
	}
	subjects, err := blls.UrbsAcAcl.FindUsersByObjects(ctx, objects)
	if err != nil {
		return nil, err
	}
	for _, setting := range ress.Result {
		setting.Users = subjects[args.Product+setting.Module+setting.Name]
	}
	return ress, nil
}

// List ...
func (a *Setting) List(ctx context.Context, args *tpl.ProductModuleURL) (*tpl.SettingsInfoRes, error) {
	ress, err := a.services.UrbsSetting.SettingList(ctx, args)
	if err != nil {
		return nil, err
	}
	objects := make([]string, len(ress.Result))
	for i, setting := range ress.Result {
		objects[i] = args.Product + args.Module + setting.Name
	}
	subjects, err := blls.UrbsAcAcl.FindUsersByObjects(ctx, objects)
	if err != nil {
		return nil, err
	}
	for _, setting := range ress.Result {
		setting.Users = subjects[args.Product+args.Module+setting.Name]
	}
	return ress, nil
}

// Get ...
func (a *Setting) Get(ctx context.Context, args *tpl.ProductModuleSettingURL) (*tpl.SettingInfoRes, error) {
	return a.services.UrbsSetting.SettingGet(ctx, args)
}

// ListGroups ...
func (a *Setting) ListGroups(ctx context.Context, args *tpl.ProductModuleSettingURL) (*tpl.SettingGroupsInfoRes, error) {
	return a.services.UrbsSetting.SettingListGroups(ctx, args)
}

// ListUsers ...
func (a *Setting) ListUsers(ctx context.Context, args *tpl.ProductModuleSettingURL) (*tpl.SettingUsersInfoRes, error) {
	return a.services.UrbsSetting.SettingListUsers(ctx, args)
}

// Create 创建指定产品功能模块配置项
func (a *Setting) Create(ctx context.Context, args *tpl.ProductModuleURL, body *tpl.SettingBody) (*tpl.SettingInfoRes, error) {
	object := args.Product + args.Module + body.Name
	err := blls.UrbsAcAcl.AddDefaultPermission(ctx, body.Uids, object)
	if err != nil {
		return nil, err
	}
	res, err := a.services.UrbsSetting.SettingCreate(ctx, args, body)
	if err != nil {
		return nil, err
	}
	res.Result.Users, err = blls.UrbsAcAcl.FindUsersByObject(ctx, object)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Update 更新指定产品功能模块配置项
func (a *Setting) Update(ctx context.Context, args *tpl.ProductModuleSettingURL, body *tpl.SettingUpdateBody) (*tpl.SettingInfoRes, error) {
	object := args.Product + args.Module + args.Setting
	err := blls.UrbsAcAcl.Update(ctx, body.UidsBody, object)
	if err != nil {
		return nil, err
	}
	res, err := a.services.UrbsSetting.SettingUpdate(ctx, args, body)
	if err != nil {
		return nil, err
	}
	res.Result.Users, err = blls.UrbsAcAcl.FindUsersByObject(ctx, object)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Offline 下线指定产品功能模块配置项
func (a *Setting) Offline(ctx context.Context, args *tpl.ProductModuleSettingURL) (*tpl.BoolRes, error) {
	return a.services.UrbsSetting.SettingOffline(ctx, args)
}

// Assign 批量为用户或群组设置产品功能模块配置项
func (a *Setting) Assign(ctx context.Context, args *tpl.ProductModuleSettingURL, body *tpl.UsersGroupsBody) (*tpl.SettingReleaseInfoRes, error) {
	AddUserAndOrg(ctx, body.Users, body.Groups)
	res, err := a.services.UrbsSetting.SettingAssign(ctx, args, body)
	if err != nil {
		return nil, err
	}
	object := args.Product + args.Module + args.Setting
	logContent := &dto.OperationLogContent{
		Users:   body.Users,
		Groups:  body.Groups,
		Desc:    body.Desc,
		Value:   body.Value,
		Release: res.Result.Release,
	}
	err = blls.OperationLog.Add(ctx, object, actionCreate, logContent)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Recall ...
func (a *Setting) Recall(ctx context.Context, args *tpl.ProductModuleSettingURL, body *tpl.RecallBody) (*tpl.BoolRes, error) {
	logID := service.HIDToID(body.HID, "log")
	log, err := daos.OperationLog.FindOneByID(ctx, logID)
	if err != nil {
		return nil, err
	}
	release := getRelease(log.Content)
	if release < 1 {
		return nil, gear.ErrBadRequest.WithMsgf("invalid release %d", release)
	}
	body.Release = release
	recallRes, err := a.services.UrbsSetting.SettingRecall(ctx, args, body)
	if err != nil {
		return nil, err
	}
	err = daos.OperationLog.DeleteByObject(ctx, logID)
	if err != nil {
		return nil, err
	}
	logger.Info(ctx, "labelRecall", "operator", util.GetUid(ctx), "log", log.String)
	return recallRes, nil
}

// ListRules ...
func (a *Setting) ListRules(ctx context.Context, args *tpl.ProductModuleSettingURL) (*tpl.SettingRulesInfoRes, error) {
	return a.services.UrbsSetting.SettingListRule(ctx, args)
}

// CreateRule ...
func (a *Setting) CreateRule(ctx context.Context, args *tpl.ProductModuleSettingURL, body *tpl.SettingRuleBody) (*tpl.SettingRuleInfoRes, error) {
	object := args.Product + args.Module + args.Setting
	logContent := &dto.OperationLogContent{
		Desc:    body.Desc,
		Percent: body.Rule.Value,
		Value:   body.Value,
	}
	err := blls.OperationLog.Add(ctx, object, actionCreate, logContent)
	if err != nil {
		return nil, err
	}
	return a.services.UrbsSetting.SettingCreateRule(ctx, args, body)
}

// UpdateRule ...
func (a *Setting) UpdateRule(ctx context.Context, args *tpl.ProductModuleSettingHIDURL, body *tpl.SettingRuleBody) (*tpl.SettingRuleInfoRes, error) {
	object := args.Product + args.Module + args.Setting
	logContent := &dto.OperationLogContent{
		Desc:    body.Desc,
		Percent: body.Rule.Value,
		Value:   body.Value,
	}
	err := blls.OperationLog.Add(ctx, object, actionUpdate, logContent)
	if err != nil {
		return nil, err
	}
	return a.services.UrbsSetting.SettingUpdateRule(ctx, args, body)
}

// DeleteRule ...
func (a *Setting) DeleteRule(ctx context.Context, args *tpl.ProductModuleSettingHIDURL) (*tpl.BoolRes, error) {
	return a.services.UrbsSetting.SettingDeleteRule(ctx, args)
}

// DeleteUser ...
func (a *Setting) DeleteUser(ctx context.Context, args *tpl.ProductModuleSettingUIDURL) (*tpl.BoolRes, error) {
	return a.services.UrbsSetting.SettingDeleteUser(ctx, args)
}

// DeleteGroup ...
func (a *Setting) DeleteGroup(ctx context.Context, args *tpl.ProductModuleSettingUIDURL) (*tpl.BoolRes, error) {
	return a.services.UrbsSetting.SettingDeleteGroup(ctx, args)
}

// RollbackGroupSetting ...
func (a *Setting) RollbackGroupSetting(ctx context.Context, args *tpl.ProductModuleSettingUIDURL) (*tpl.BoolRes, error) {
	return a.services.UrbsSetting.SettingRollbackGroupSetting(ctx, args)
}

// RollbackUserSetting ...
func (a *Setting) RollbackUserSetting(ctx context.Context, args *tpl.ProductModuleSettingUIDURL) (*tpl.BoolRes, error) {
	return a.services.UrbsSetting.SettingRollbackUserSetting(ctx, args)
}
