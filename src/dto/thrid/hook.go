package thrid

// HookSendReq ...
type HookSendReq struct {
	Event   string   `json:"event"`
	Users   []string `json:"users"`
	Content string   `json:"content"`
}
