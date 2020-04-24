package tpl

import (
	"time"

	"github.com/teambition/gear"
)

// UrbsSettingUser ...
type UrbsSettingUser struct {
	CreatedAt time.Time `json:"createdAt"`
	UID       string    `json:"uid"`
	ActiveAt  int64     `json:"activeAt"`
	Labels    string    `json:"labels"`
}

// UsersBody ...
type UsersBody struct {
	Users []string `json:"users"`
}

// Validate 实现 gear.BodyTemplate。
func (t *UsersBody) Validate() error {
	if len(t.Users) == 0 {
		return gear.ErrBadRequest.WithMsg("users emtpy")
	}
	for _, uid := range t.Users {
		if !validIDReg.MatchString(uid) {
			return gear.ErrBadRequest.WithMsgf("invalid user: %s", uid)
		}
	}
	return nil
}

// UsersRes ...
type UsersRes struct {
	SuccessResponseType
	Result []UrbsSettingUser `json:"result"`
}

// UserRes ...
type UserRes struct {
	SuccessResponseType
	Result UrbsSettingUser `json:"result"`
}
