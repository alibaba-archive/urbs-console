package api

import (
	"os"
	"testing"

	"github.com/teambition/gear"
)

var (
	urbsSettingUrl string
)

func TestMain(m *testing.M) {

	os.Exit(m.Run())
}

type TestTools struct {
	App  *gear.App
	Host string
}

func SetUpTestTools() (tt *TestTools, cleanup func()) {
	tt = &TestTools{}
	tt.App = NewApp()
	srv := tt.App.Start()
	tt.Host = "http://" + srv.Addr().String()

	return tt, func() {
		srv.Close()
	}
}
