package api

import (
	"errors"
	"net/http"
	"os"
	"testing"

	"github.com/teambition/gear"
	"github.com/teambition/urbs-console/src/dto/thrid"
)

var (
	testHost string
	mockUser = "827004504e8c60"
)

func TestMain(m *testing.M) {
	app := NewApp()
	srv := app.Start()
	testHost = "http://" + srv.Addr().String()

	services.UserAuth = &MockUserAuthInterface{}
	os.Exit(m.Run())
}

func userTokenHeader() http.Header {
	header := http.Header{}
	header.Set("Authorization", "Bearer user")
	return header
}

// MockUserAuthInterface ...
type MockUserAuthInterface struct {
}

// Verify ...
func (a *MockUserAuthInterface) Verify(ctx *gear.Context, body *thrid.UserVerifyReq) (string, error) {
	if body.Token == "user" {
		return mockUser, nil
	}
	return "", errors.New("invalid token")
}
