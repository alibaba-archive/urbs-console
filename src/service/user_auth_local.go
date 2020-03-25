package service

import (
	"time"

	"github.com/teambition/gear"
	auth "github.com/teambition/gear-auth"
	authjwt "github.com/teambition/gear-auth/jwt"
	"github.com/teambition/gear/logging"
	"github.com/teambition/urbs-console/src/conf"
	"github.com/teambition/urbs-console/src/logger"
)

func init() {
	keys := conf.Config.UserAuth.Keys
	if len(keys) > 0 {
		Auther = auth.New(authjwt.StrToKeys(keys...)...)
		Auther.JWT().SetExpiresIn(time.Minute * 10)
	} else {
		logger.Default.Warningf("`user_auth.keys` is empty, Auth middleware will not be executed.")
	}
}

// UserAuthLocal ...
type UserAuthLocal struct {
}

// Auther ...
var Auther *auth.Auth

// Verify ...
func (a *UserAuthLocal) Verify(ctx *gear.Context) error {
	if Auther != nil {
		claims, err := Auther.FromCtx(ctx)
		if err != nil {
			return err
		}
		if sub, ok := claims.Subject(); ok {
			logging.SetTo(ctx, "jwt_sub", sub)
		}
		if jti, ok := claims.JWTID(); ok {
			logging.SetTo(ctx, "jwt_id", jti)
		}
	}
	return nil
}
