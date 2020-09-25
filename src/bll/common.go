package bll

import (
	"github.com/teambition/urbs-console/src/dto"
	"github.com/teambition/urbs-console/src/tpl"
)

type operationLogAdd struct {
	Object  string                   `json:"object"`
	Content *dto.OperationLogContent `json:"content"`
	Action  string                   `json:"Action"`
}

type settingRecallReq struct {
	Args *tpl.ProductModuleSettingURL `json:"args"`
	Body *tpl.RecallBody              `json:"body"`
}

type labelRecallReq struct {
	Args *tpl.ProductLabelURL `json:"args"`
	Body *tpl.RecallBody      `json:"body"`
}

// MatchClient ...
func MatchClient(clients []string, client string) bool {
	if len(clients) == 0 || client == "" {
		return true
	}
	for _, c := range clients {
		if c == client {
			return true
		}
	}
	return false
}

// MatchChannel ...
func MatchChannel(channels []string, channel string) bool {
	if len(channels) == 0 || channel == "" {
		return true
	}
	for _, c := range channels {
		if c == channel {
			return true
		}
	}
	return false
}
