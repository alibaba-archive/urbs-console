package tpl

import "regexp"

var validIDReg = regexp.MustCompile(`^[0-9A-Za-z._=-]{3,63}$`)

// Should be subset of DNS-1035 label
// https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#dns-label-names
var validLabelReg = regexp.MustCompile(`^[0-9a-z][0-9a-z-]{0,61}[0-9a-z]$`)

// SuccessResponseType 定义了标准的 API 接口成功时返回数据模型
type SuccessResponseType struct {
	TotalSize     int         `json:"totalSize,omitempty"`
	NextPageToken string      `json:"nextPageToken,omitempty"`
	Result        interface{} `json:"result"`
}

// BoolRes ...
type BoolRes struct {
	SuccessResponseType
	Result bool `json:"result"`
}
