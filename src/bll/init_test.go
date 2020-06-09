package bll

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/teambition/urbs-console/src/dao"
	"github.com/teambition/urbs-console/src/service"
	"github.com/teambition/urbs-console/src/util"
)

var (
	testDaos *dao.Daos
	testDB   *service.SQL
	testBlls *Blls
)

func TestMain(m *testing.M) {
	testDB = service.NewDB()
	testDaos = dao.NewDaos(service.NewDB())

	testBlls = NewBlls(service.NewServices(testDB), testDaos)
	os.Exit(m.Run())
}

type TestTools struct {
	Require *require.Assertions
}

func SetUpTestTools(r *require.Assertions) *TestTools {
	tt := &TestTools{}
	tt.Require = r
	return tt
}

func getUidContext(uid ...string) context.Context {
	if len(uid) > 0 {
		return context.WithValue(context.Background(), util.UidKey{}, uid[0])
	}
	return context.WithValue(context.Background(), util.UidKey{}, "123")
}
