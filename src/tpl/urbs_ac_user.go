package tpl

import (
	"github.com/teambition/gear"
	"github.com/teambition/urbs-console/src/schema"
)

// UrbsAcUsersBody ...
type UrbsAcUsersBody struct {
	Users []*UrbsAcUserBody `json:"users"`
}

// Validate 实现 gear.BodyTemplate。
func (t *UrbsAcUsersBody) Validate() error {
	if len(t.Users) == 0 {
		return gear.ErrBadRequest.WithMsg("empty users")
	}
	return nil
}

// UrbsAcUserBody ...
type UrbsAcUserBody struct {
	Uid  string `json:"uid"`
	Name string `json:"name"`
}

// UrbsAcUserListRes ...
type UrbsAcUserListRes struct {
	SuccessResponseType
	Result []*schema.UrbsAcUser `json:"result"` // 空数组也保留
}

// UrbsAcUserUrl ...
type UrbsAcUserUrl struct {
	Key string `json:"key" query:"key"`
}

// Validate 实现 gear.BodyTemplate。
func (t *UrbsAcUserUrl) Validate() error {
	if len(t.Key) == 0 || len(t.Key) > 63 {
		return gear.ErrBadRequest.WithMsg("invalid key")
	}
	return nil
}
