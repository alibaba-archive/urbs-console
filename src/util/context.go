package util

import (
	"context"
)

// UidKey ...
type UidKey struct{}

//GetUid ...
func GetUid(ctx context.Context) string {
	uid, ok := ctx.Value(UidKey{}).(string)
	if !ok {
		panic("invalid uid")
	}
	return uid
}
