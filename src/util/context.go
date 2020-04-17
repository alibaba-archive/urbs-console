package util

import (
	"context"
)

// UidKey ...
type UidKey struct{}

//GetUid ...
func GetUid(ctx context.Context) string {
	uid, _ := ctx.Value(UidKey{}).(string)
	return uid
}
