package service

import (
	"net/http"

	"github.com/teambition/urbs-console/src/conf"
	"github.com/teambition/urbs-console/src/util"
)

func init() {
	util.DigProvide(NewServices)
}

// Services ...
type Services struct {
	UrbsSetting UrbsSettingInterface

	UserAuth    UserAuth
	GroupMember *GroupMember
}

// NewServices ...
func NewServices() *Services {
	s := &Services{
		GroupMember: &GroupMember{},
		UrbsSetting: &UrbsSetting{},
	}
	if conf.Config.UserAuth.UserAuthThrid.URL == "" {
		s.UserAuth = &UserAuthLocal{}
	} else {
		s.UserAuth = &UserAuthThrid{}
	}
	return s
}

// UrbsSettingHeader ...
func UrbsSettingHeader() http.Header {
	header := http.Header{}
	return header
}
