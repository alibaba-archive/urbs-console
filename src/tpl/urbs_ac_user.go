package tpl

import "github.com/teambition/gear"

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
