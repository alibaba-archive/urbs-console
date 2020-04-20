package tpl

// UrbsAcAclAddReq ...
type UrbsAcAclAddReq struct {
	Uid string `json:"uid"`

	Product string `json:"product"`
	Label   string `json:"label"`
	Module  string `json:"module"`
	Setting string `json:"setting"`

	Permission string `json:"permission"`
}

// Validate 实现 gear.BodyTemplate。
func (t *UrbsAcAclAddReq) Validate() error {
	return nil
}
