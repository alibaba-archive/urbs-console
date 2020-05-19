package tpl

import (
	"time"
)

// LabelRuleBody ...
type LabelRuleBody struct {
	PercentRule
	Desc string `json:"desc"` // 操作说明
}

// LabelRuleInfo ...
type LabelRuleInfo struct {
	ID        int64       `json:"-"`
	HID       string      `json:"hid"`
	LabelHID  string      `json:"labelHID"`
	Kind      string      `json:"kind"`
	Rule      interface{} `json:"rule"`
	Release   int64       `json:"release"`
	CreatedAt time.Time   `json:"createdAt"`
	UpdatedAt time.Time   `json:"updatedAt"`
}

// LabelRulesInfoRes ...
type LabelRulesInfoRes struct {
	SuccessResponseType
	Result []LabelRuleInfo `json:"result"` // 空数组也保留
}

// LabelRuleInfoRes ...
type LabelRuleInfoRes struct {
	SuccessResponseType
	Result LabelRuleInfo `json:"result"` // 空数组也保留
}
