package bll

import (
	"context"
	"time"

	"github.com/teambition/urbs-console/src/dao"
	"github.com/teambition/urbs-console/src/dto/urbssetting"
	"github.com/teambition/urbs-console/src/logger"
	"github.com/teambition/urbs-console/src/service"
	"github.com/teambition/urbs-console/src/tpl"
)

// Group ...
type Group struct {
	services *service.Services
	daos     *dao.Daos
}

// ListLables ...
func (a *Group) ListLables(ctx context.Context, args *tpl.UIDPaginationURL) (*tpl.MyLabelsRes, error) {
	return a.services.UrbsSetting.GroupListLables(ctx, args)
}

// List ...
func (a *Group) List(ctx context.Context, args *tpl.GroupsURL) (*tpl.GroupsRes, error) {
	return a.services.UrbsSetting.GroupList(ctx, args)
}

// ListSettings ...
func (a *Group) ListSettings(ctx context.Context, args *tpl.UIDPaginationURL) (*tpl.MySettingsRes, error) {
	return a.services.UrbsSetting.GroupListSettings(ctx, args)
}

// CheckExists ...
func (a *Group) CheckExists(ctx context.Context, uid string) (*tpl.BoolRes, error) {
	return a.services.UrbsSetting.GroupCheckExists(ctx, uid)
}

// BatchAdd ...
func (a *Group) BatchAdd(ctx context.Context, groups []tpl.GroupBody) error {
	_, err := a.services.UrbsSetting.GroupBatchAdd(ctx, groups)
	if err != nil {
		return err
	}
	for i := range groups {
		go func(group *tpl.GroupBody) {
			err := a.daos.UrbsLock.Lock(ctx, group.Kind+group.UID, 30*time.Minute)
			if err == nil {
				a.BatchAddMember(ctx, group.UID)
				a.daos.UrbsLock.Unlock(ctx, group.Kind+group.UID)
			}
		}(&groups[i])
	}
	return nil
}

// BatchAddMember ...
func (a *Group) BatchAddMember(ctx context.Context, uid string) error {
	pageSize := 1000
	count := 0
	now := time.Now().Unix()
	// 更新同步时间
	groupUpdateBody := new(urbssetting.GroupUpdateBody)
	groupUpdateBody.SyncAt = &now
	_, err := a.services.UrbsSetting.GroupUpdate(ctx, uid, groupUpdateBody)
	if err != nil {
		logger.Err(ctx, "groupUpdate", "error", err.Error())
		return err
	}
	nextPageToken := ""
	// 同步成员
	for {
		var resp *service.ListGroupMembersResp
		resp, err = a.services.GroupMember.List(ctx, uid, nextPageToken, pageSize)
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
	if err != nil { // 同步错误，不删除旧的成员
		return nil
	}
	// 删除旧的成员
	args := new(tpl.GroupMembersURL)
	args.UID = uid
	args.SyncLt = now
	_, err = a.services.UrbsSetting.GroupRemoveMembers(ctx, args)
	if err != nil {
		logger.Err(ctx, "groupRemoveMembers", "error", err.Error())
	} else {
		logger.Info(ctx, "batchAddMember", "count", count, "uid", uid)
	}
	return nil
}

// Update ...
func (a *Group) Update(ctx context.Context, uid string, body *tpl.GroupUpdateBody) (*tpl.GroupRes, error) {
	b := &urbssetting.GroupUpdateBody{
		Desc: body.Desc,
	}
	return a.services.UrbsSetting.GroupUpdate(ctx, uid, b)
}

// Delete ...
func (a *Group) Delete(ctx context.Context, uid string) (*tpl.BoolRes, error) {
	return a.services.UrbsSetting.GroupDelete(ctx, uid)
}

// ListMembers ...
func (a *Group) ListMembers(ctx context.Context, args *tpl.UIDPaginationURL) (*tpl.GroupMembersRes, error) {
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
