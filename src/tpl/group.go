package tpl

import (
	"time"

	"github.com/teambition/gear"
)

// Group ...
type Group struct {
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	SyncAt    int64     `json:"syncAt"`
	UID       string    `json:"uid"`
	Kind      string    `json:"kind"`
	Desc      string    `json:"desc"`
	Status    int64     `json:"status"`
}

// GroupsBody ...
type GroupsBody struct {
	Groups []GroupBody `json:"groups"`
}

// GroupBody ...
type GroupBody struct {
	UID  string `json:"uid"`
	Kind string `json:"kind"`
	Desc string `json:"desc"`
}

// Validate 实现 gear.BodyTemplate。
func (t *GroupsBody) Validate() error {
	if len(t.Groups) == 0 {
		return gear.ErrBadRequest.WithMsg("groups emtpy")
	}
	for _, g := range t.Groups {
		if !validIDReg.MatchString(g.UID) {
			return gear.ErrBadRequest.WithMsgf("invalid group uid: %s", g.UID)
		}
		if !validLabelReg.MatchString(g.Kind) {
			return gear.ErrBadRequest.WithMsgf("invalid group kind: %s", g.Kind)
		}
		if len(g.Desc) > 1022 {
			return gear.ErrBadRequest.WithMsgf("desc too long: %d", len(g.Desc))
		}
	}
	return nil
}

// GroupsURL ...
type GroupsURL struct {
	Pagination
	Kind string `json:"kind" query:"kind"`
}

// GroupUpdateBody ...
type GroupUpdateBody struct {
	Desc *string `json:"desc"`
}

// Validate 实现 gear.BodyTemplate。
func (t *GroupUpdateBody) Validate() error {
	if t.Desc == nil {
		return gear.ErrBadRequest.WithMsgf("desc or kind or syncAt required")
	}
	if len(*t.Desc) > 1022 {
		return gear.ErrBadRequest.WithMsgf("desc too long: %d", len(*t.Desc))
	}
	return nil
}

// GroupMembersURL ...
type GroupMembersURL struct {
	Kind   string `json:"kind" query:"kind"`
	UID    string `json:"uid" param:"uid"`
	User   string `json:"user" query:"user"`     // 根据用户 uid 删除一个成员
	SyncLt int64  `json:"syncLt" query:"syncLt"` // 或根据 syncLt 删除同步时间小于指定值的所有成员
}

// Validate 实现 gear.BodyTemplate。
func (t *GroupMembersURL) Validate() error {
	if !validIDReg.MatchString(t.UID) {
		return gear.ErrBadRequest.WithMsgf("invalid group uid: %s", t.UID)
	}

	if t.User != "" {
		if !validIDReg.MatchString(t.User) {
			return gear.ErrBadRequest.WithMsgf("invalid user uid: %s", t.User)
		}
	} else if t.SyncLt != 0 {
		if t.SyncLt < 0 || t.SyncLt > (time.Now().UTC().Unix()+3600) {
			// 较大的 SyncLt 可以删除整个群组成员！+3600 是防止把毫秒当秒用
			return gear.ErrBadRequest.WithMsgf("invalid syncLt: %d", t.SyncLt)
		}
	} else {
		return gear.ErrBadRequest.WithMsg("user or syncLt required")
	}
	return nil
}

// GroupsRes ...
type GroupsRes struct {
	SuccessResponseType
	Result []Group `json:"result"`
}

// GroupMember ...
type GroupMember struct {
	User      string    `json:"user"`
	CreatedAt time.Time `json:"createdAt"`
	SyncAt    int64     `json:"syncAt"` // 归属关系同步时间戳，1970 以来的秒数，应该与 group.syncAt 相等
}

// GroupMembersRes ...
type GroupMembersRes struct {
	SuccessResponseType
	Result []GroupMember `json:"result"`
}

// GroupRes ...
type GroupRes struct {
	SuccessResponseType
	Result Group `json:"result"`
}

// GroupURL ...
type GroupURL struct {
	Kind string `json:"kind" query:"kind"`
	UID  string `json:"uid" param:"uid"`
}

// Validate 实现 gear.BodyTemplate。
func (t *GroupURL) Validate() error {
	if !validIDReg.MatchString(t.UID) {
		return gear.ErrBadRequest.WithMsgf("invalid uid: %s", t.UID)
	}
	return nil
}

// GroupPaginationURL ...
type GroupPaginationURL struct {
	Pagination
	UID  string `json:"uid" param:"uid"`
	Kind string `json:"kind" query:"kind"`
}

// Validate 实现 gear.BodyTemplate。
func (t *GroupPaginationURL) Validate() error {
	if !validIDReg.MatchString(t.UID) {
		return gear.ErrBadRequest.WithMsgf("invalid uid: %s", t.UID)
	}
	if err := t.Pagination.Validate(); err != nil {
		return err
	}
	return nil
}
