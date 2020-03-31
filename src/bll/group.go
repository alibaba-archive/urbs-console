package bll

import (
	"context"
	"time"

	"github.com/teambition/urbs-console/src/dto/urbssetting"
	"github.com/teambition/urbs-console/src/logger"
	"github.com/teambition/urbs-console/src/service"
	"github.com/teambition/urbs-console/src/tpl"
)

// Group ...
type Group struct {
	services *service.Services
}

// ListLables ...
func (a *Group) ListLables(ctx context.Context, args *tpl.UIDPaginationURL) (*urbssetting.LabelsInfoRes, error) {
	return a.services.UrbsSetting.GroupListLables(ctx, args)
}

// List ...
func (a *Group) List(ctx context.Context, args *tpl.GroupsURL) (*urbssetting.GroupsRes, error) {
	return a.services.UrbsSetting.GroupList(ctx, args)
}

// ListSettings ...
func (a *Group) ListSettings(ctx context.Context, args *tpl.UIDProductURL) (*urbssetting.MySettingsRes, error) {
	return a.services.UrbsSetting.GroupListSettings(ctx, args)
}

// CheckExists ...
func (a *Group) CheckExists(ctx context.Context, uid string) (*tpl.BoolRes, error) {
	return a.services.UrbsSetting.GroupCheckExists(ctx, uid)
}

// BatchAdd ...
func (a *Group) BatchAdd(ctx context.Context, groups []*tpl.GroupBody) error {
	_, err := a.services.UrbsSetting.GroupBatchAdd(ctx, groups)
	if err != nil {
		return err
	}
	for _, g := range groups {
		go a.BatchAddMember(ctx, g.UID)
	}
	return nil
}

// BatchAddMember ...
func (a *Group) BatchAddMember(ctx context.Context, uid string) error {
	pageSize := 1000
	count := 0
	now := time.Now().Unix()
	// 更新同步时间
	groupUpdateBody := new(tpl.GroupUpdateBody)
	groupUpdateBody.SyncAt = &now
	_, err := a.services.UrbsSetting.GroupUpdate(ctx, uid, groupUpdateBody)
	if err != nil {
		logger.Err(ctx, "groupUpdate", "error", err.Error())
		return err
	}
	nextPageToken := ""
	// 同步成员
	for {
		resp, err := a.services.GroupMember.List(uid, nextPageToken, pageSize)
		if err != nil {
			logger.Err(ctx, err.Error(), "uid", uid)
			break
		}
		nextPageToken = resp.NextPageToken
		count += len(resp.Members)

		users := make([]string, len(resp.Members))
		for i, r := range resp.Members {
			users[i] = r.UID
		}
		_, err = a.services.UrbsSetting.GroupBatchAddMembers(ctx, uid, users)
		if err != nil {
			logger.Err(ctx, err.Error(), "uid", uid)
		}
		if len(resp.Members) >= pageSize {
			continue
		}
		break
	}
	// 删除旧的成员
	args := new(tpl.GroupMembersURL)
	args.UID = uid
	args.SyncLt = now
	_, err = a.services.UrbsSetting.GroupRemoveMembers(ctx, args)
	if err != nil {
		logger.Err(ctx, "groupRemoveMembers", "error", err.Error())
	}
	logger.Info(ctx, "batchAddMember", "count", count, "uid", uid)
	return nil
}

// Update ...
func (a *Group) Update(ctx context.Context, uid string, body *tpl.GroupUpdateBody) (*urbssetting.GroupRes, error) {
	return a.services.UrbsSetting.GroupUpdate(ctx, uid, body)
}

// Delete ...
func (a *Group) Delete(ctx context.Context, uid string) (*tpl.BoolRes, error) {
	return a.services.UrbsSetting.GroupDelete(ctx, uid)
}

// ListMembers ...
func (a *Group) ListMembers(ctx context.Context, args *tpl.UIDPaginationURL) (*urbssetting.GroupMembersRes, error) {
	return a.services.UrbsSetting.GroupListMembers(ctx, args)
}

// BatchAddMembers 批量给群组添加成员，如果用户未加入系统，则会自动加入
func (a *Group) BatchAddMembers(ctx context.Context, groupId string, users []string) (*tpl.BoolRes, error) {
	return a.services.UrbsSetting.GroupBatchAddMembers(ctx, groupId, users)
}

// RemoveMembers ...
func (a *Group) RemoveMembers(ctx context.Context, args *tpl.GroupMembersURL) (*tpl.BoolRes, error) {
	return a.services.UrbsSetting.GroupRemoveMembers(ctx, args)
}

// RemoveLable ...
func (a *Group) RemoveLable(ctx context.Context, args *tpl.UIDHIDURL) (*tpl.BoolRes, error) {
	return a.services.UrbsSetting.GroupRemoveLable(ctx, args)
}

// RollbackSetting 回滚指定群组的指定配置项
func (a *Group) RollbackSetting(ctx context.Context, args *tpl.UIDHIDURL) (*tpl.BoolRes, error) {
	return a.services.UrbsSetting.GroupRollbackSetting(ctx, args)
}

// RemoveSetting 删除指定群组的指定配置项
func (a *Group) RemoveSetting(ctx context.Context, args *tpl.UIDHIDURL) (*tpl.BoolRes, error) {
	return a.services.UrbsSetting.GroupRemoveSetting(ctx, args)
}
