package bll

import "github.com/teambition/urbs-console/src/util"

func init() {
	util.DigProvide(NewBlls)
}

var (
	blls *Blls
)
