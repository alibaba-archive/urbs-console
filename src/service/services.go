package service

// Services ...
type Services struct {
	UrbsSetting UrbsSettingInterface
	UserAuth    UserAuthInterface
	GroupMember GroupMemberInterface
	Hook        HookInterface
}

// NewServices ...
func NewServices() *Services {
	s := &Services{
		GroupMember: &GroupMember{},
		UrbsSetting: &UrbsSetting{},
		UserAuth:    &UserAuth{},
		Hook:        &Hook{},
	}
	return s
}
