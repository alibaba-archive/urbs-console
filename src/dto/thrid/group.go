package thrid

// ListGroupMembersResp ...
type ListGroupMembersResp struct {
	Members       []Member `json:"result"`
	NextPageToken string   `json:"nextPageToken"`
}

// Member ....
type Member struct {
	UID string `json:"uid"`
}
