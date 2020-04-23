package tpl

import (
	"github.com/teambition/gear"
	"github.com/teambition/urbs-console/src/schema"
)

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
	Result []schema.User `json:"result"`
}

// UserRes ...
type UserRes struct {
	SuccessResponseType
	Result schema.User `json:"result"`
}
