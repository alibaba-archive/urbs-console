package thrid

// HookSendReq ...
type HookSendReq struct {
	Event   string   `json:"event"`
	Users   []string `json:"users,omitempty"`
	Content string   `json:"content"`
}
