package service

import (
	"context"

	"github.com/teambition/urbs-console/src/dto/urbssetting"
)

// UrbsSettingInterface ....
type UrbsSettingInterface interface {
	UserBatchAdd(ctx context.Context, users []string) (*urbssetting.BoolRes, error)
}
