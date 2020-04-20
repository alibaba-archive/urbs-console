package urbssetting

import "time"

// MySettingsRes ...
type MySettingsRes struct {
	SuccessResponseType
	Result []MySetting `json:"result"` // 空数组也保留
}

// MySetting ...
type MySetting struct {
	ID        int64     `json:"-"`
	HID       string    `json:"hid"`
	Module    string    `json:"module"`
	Name      string    `json:"name"`
	Value     string    `json:"value"`
	LastValue string    `json:"last_value"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SettingInfo ...
type SettingInfo struct {
	ID        int64      `json:"-"`
	HID       string     `json:"hid"`
	Product   string     `json:"product"`
	Module    string     `json:"module"`
	Name      string     `json:"name"`
	Desc      string     `json:"desc"`
	Channels  []string   `json:"channels"`
	Clients   []string   `json:"clients"`
	Values    []string   `json:"values"`
	Status    int64      `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	OfflineAt *time.Time `json:"offline_at"`
}
