package tpl

import (
	"encoding/json"
	"fmt"
	"time"
)

// PercentRule ...
type PercentRule struct {
	Kind string `json:"kind"`
	Rule struct {
		Value int `json:"value"`
	} `json:"rule"`
}

// Validate ...
func (r *PercentRule) Validate() error {
	if r.Kind != "userPercent" {
		return fmt.Errorf("invalid kind: %s", r.Kind)
	}
	if r.Rule.Value < 0 || r.Rule.Value > 100 {
		return fmt.Errorf("invalid percent rule value: %d", r.Rule.Value)
	}
	return nil
}

// ToRule ...
func (r *PercentRule) ToRule() string {
	if b, err := json.Marshal(r.Rule); err == nil {
		return string(b)
	}
	return ""
}

// SettingRuleBody ...
type SettingRuleBody struct {
	PercentRule
	Value string `json:"value"`
	Desc  string `json:"desc"` // 操作说明
}

// SettingRuleInfo ...
type SettingRuleInfo struct {
	ID         int64       `json:"-"`
	HID        string      `json:"hid"`
	SettingHID string      `json:"settingHID"`
	Kind       string      `json:"kind"`
	Rule       interface{} `json:"rule"`
	Value      string      `json:"value"`
	Release    int64       `json:"release"`
	CreatedAt  time.Time   `json:"createdAt"`
	UpdatedAt  time.Time   `json:"updatedAt"`
}

// SettingRulesInfoRes ...
type SettingRulesInfoRes struct {
	SuccessResponseType
	Result []SettingRuleInfo `json:"result"` // 空数组也保留
}

// SettingRuleInfoRes ...
type SettingRuleInfoRes struct {
	SuccessResponseType
	Result SettingRuleInfo `json:"result"` // 空数组也保留
}
