package urbssetting

// UsersGroupsBody ...
type UsersGroupsBody struct {
	Users  []string        `json:"users"`
	Groups []*GroupKindUID `json:"groups"`
	Value  string          `json:"value"`
}

// GroupKindUID ...
type GroupKindUID struct {
	Kind string `json:"kind"`
	UID  string `json:"uid"`
}
