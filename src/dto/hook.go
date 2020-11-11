package dto

import (
	"encoding/json"
)

// HookCleanup ...
type HookCleanup struct {
	Product string `json:"product,omitempty"`

	Module  string `json:"module,omitempty" `
	Setting string `json:"setting,omitempty"`

	Label string `json:"label,omitempty"`

	Operator string `json:"operator"` // 操作人
}

// Marshal ...
func (a *HookCleanup) Marshal() string {
	b, err := json.Marshal(a)
	if err != nil {
		panic(err.Error())
	}
	return string(b)
}

// HookRule ...
type HookRule struct {
	Product string `json:"product,omitempty"`

	Module  string `json:"module,omitempty" `
	Setting string `json:"setting,omitempty"`

	Label string `json:"label,omitempty"`

	Kind    string `json:"kind,omitempty"`
	Percent *int   `json:"percent,omitempty"`
	Desc    string `json:"desc,omitempty"`

	Operator string `json:"operator"` // 操作人
}

// Marshal ...
func (a *HookRule) Marshal() string {
	b, err := json.Marshal(a)
	if err != nil {
		panic(err.Error())
	}
	return string(b)
}
