package urbssetting

// GroupUpdateBody ...
type GroupUpdateBody struct {
	Desc   *string `json:"desc"`
	SyncAt *int64  `json:"syncAt"`
}
