package tpl

import "time"

// SettingGroupInfo ...
type SettingGroupInfo struct {
	SettingHID string    `json:"settingHID"`
	AssignedAt time.Time `json:"assignedAt"`
	Release    int64     `json:"release"`
	Group      string    `json:"Group"`
	Kind       string    `json:"kind"`
	Desc       string    `json:"desc"`
	Status     int64     `json:"status"`
	Value      string    `json:"value"`
	LastValue  string    `json:"lastValue"`
}

// SettingGroupsInfoRes ...
type SettingGroupsInfoRes struct {
	SuccessResponseType
	Result []*SettingGroupInfo `json:"result"`
}
