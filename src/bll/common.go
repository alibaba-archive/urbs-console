package bll

import "github.com/teambition/urbs-console/src/util"

func init() {
	util.DigProvide(NewBlls)
}

// Blls ...
type Blls struct {
	User  *User
	Group *Group
}

// NewBlls ...
func NewBlls() *Blls {
	return &Blls{
		User:  &User{},
		Group: &Group{},
	}
}
