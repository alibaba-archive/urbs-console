package service

import (
	"github.com/mushroomsir/tcc"
	"github.com/mushroomsir/tcc/lock"
	"github.com/mushroomsir/tcc/store"
)

// Services ...
type Services struct {
	UrbsSetting UrbsSettingInterface
	UserAuth    UserAuthInterface
	GroupMember GroupMemberInterface
	Hook        HookInterface
	TCC         *tcc.TCC
}

// NewServices ...
func NewServices(sql *SQL) *Services {
	s := &Services{
		GroupMember: &GroupMember{},
		UrbsSetting: &UrbsSetting{},
		UserAuth:    &UserAuth{},
		Hook:        &Hook{},
	}
	option := &tcc.Option{
		PullTaskInterval: 3,
		Store:            store.NewMysql(sql.DB),
		Lock:             lock.NewMysql(sql.DB),
	}
	s.TCC = tcc.New(option)
	return s
}
