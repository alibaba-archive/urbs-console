package service

import (
	"context"

	"github.com/mushroomsir/request"
	"github.com/teambition/urbs-console/src/conf"
	"github.com/teambition/urbs-console/src/dto/thrid"
	"github.com/teambition/urbs-console/src/logger"
	"github.com/teambition/urbs-console/src/tpl"
)

const (
	// EventSettingPublish ...
	EventSettingPublish = "setting.publish"
	// EventSettingRecall ...
	EventSettingRecall = "setting.recall"
	// EventSettingRemove ...
	EventSettingRemove = "setting.remove"
)

// HookInterface ...
type HookInterface interface {
	SendAsync(ctx context.Context, body *thrid.HookSendReq)
	Send(ctx context.Context, body *thrid.HookSendReq) error
}

// Hook ...
type Hook struct {
}

// SendAsync ...
func (a *Hook) SendAsync(ctx context.Context, body *thrid.HookSendReq) {
	go func() {
		err := a.Send(ctx, body)
		if err != nil {
			logger.Err(ctx, err.Error())
		}
	}()
}

// Send ...
func (a *Hook) Send(ctx context.Context, body *thrid.HookSendReq) error {
	result := new(tpl.BoolRes)
	resp, err := request.Post(conf.Config.Thrid.Hook.URL).Header(ThridHeader(ctx)).Body(body).Result(result).Do()
	if err := HanderResponse(resp, err); err != nil {
		return err
	}
	return nil
}
