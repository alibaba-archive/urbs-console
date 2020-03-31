package api

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/teambition/gear"
	"github.com/teambition/gear-auth/jwt"
	"github.com/teambition/urbs-console/src/conf"
)

var (
	urbsSettingUrl string
)

func TestMain(m *testing.M) {

	os.Exit(m.Run())
}

type TestTools struct {
	App     *gear.App
	Host    string
	Require *require.Assertions
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

func genHeader() http.Header {
	header := http.Header{}
	header.Set("Authorization", "Bearer "+genToken())
	return header
}

func genToken() string {
	j := jwt.New([]byte(conf.Config.Thrid.Key))
	m := make(map[string]interface{})
	m["name"] = "urbs-console"
	token, err := j.Sign(m, time.Hour)
	if err != nil {
		panic(err)
	}
	return token
}
