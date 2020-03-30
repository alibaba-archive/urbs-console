package thrid

// UserVerifyReq ...
type UserVerifyReq struct {
	Role   string `json:"role"`
	Cookie string `json:"cookie"`
	Singed string `json:"signed"`
	Token  string `json:"token"`
}
