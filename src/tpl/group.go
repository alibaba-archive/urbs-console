package tpl

import "github.com/teambition/gear"

// GroupsBody ...
type GroupsBody struct {
	Groups []GroupBody `json:"groups"`
}

// GroupBody ...
type GroupBody struct {
	UID  string `json:"uid"`
	Kind string `json:"kind"`
	Desc string `json:"desc"`
}

// Validate 实现 gear.BodyTemplate。
func (t *GroupsBody) Validate() error {
	if len(t.Groups) == 0 {
		return gear.ErrBadRequest.WithMsg("groups emtpy")
	}
	for _, g := range t.Groups {
		if !validIDReg.MatchString(g.UID) {
			return gear.ErrBadRequest.WithMsgf("invalid group uid: %s", g.UID)
		}
		if !validLabelReg.MatchString(g.Kind) {
			return gear.ErrBadRequest.WithMsgf("invalid group kind: %s", g.Kind)
		}
		if len(g.Desc) > 1022 {
			return gear.ErrBadRequest.WithMsgf("desc too long: %d", len(g.Desc))
		}
	}
	return nil
}
