package urbssetting

import (
	"regexp"
)

var validIDReg = regexp.MustCompile(`^[0-9A-Za-z._=-]{3,63}$`)

// SuccessResponseType 定义了标准的 API 接口成功时返回数据模型
type SuccessResponseType struct {
	TotalSize     int         `json:"totalSize,omitempty"`
	NextPageToken string      `json:"nextPageToken,omitempty"`
	Result        interface{} `json:"result"`
}
