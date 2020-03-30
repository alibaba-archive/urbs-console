package thrid

// GroupMembersResp ...
type GroupMembersResp struct {
	Members []*Member `json:"result"`
}

// Member ....
type Member struct {
	UID string `json:"uid"`
}
