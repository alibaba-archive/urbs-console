package service

// Services ...
type Services struct {
	UrbsSetting UrbsSettingInterface
	UserAuth    UserAuthInterface
	GroupMember GroupMemberInterface
}

// NewServices ...
func NewServices() *Services {
	s := &Services{
		GroupMember: &GroupMember{},
		UrbsSetting: &UrbsSetting{},
		UserAuth:    &UserAuth{},
	}
	return s
}
