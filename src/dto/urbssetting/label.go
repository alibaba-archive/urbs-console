package urbssetting

import "time"

// LabelBody ...
type LabelBody struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

// LabelsInfoRes ...
type LabelsInfoRes struct {
	SuccessResponseType
	Result []*LabelInfo `json:"result"` // 空数组也保留
}

// LabelInfo ...
type LabelInfo struct {
	ID        int64      `json:"-"`
	HID       string     `json:"hid"`
	Product   string     `json:"product"`
	Name      string     `json:"name"`
	Desc      string     `json:"desc"`
	Channels  []string   `json:"channels"`
	Clients   []string   `json:"clients"`
	Status    int64      `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	OfflineAt *time.Time `json:"offline_at"`
}

// LabelUpdateBody ...
type LabelUpdateBody struct {
	Desc     *string   `json:"desc"`
	Channels *[]string `json:"channels"`
	Clients  *[]string `json:"clients"`
}
