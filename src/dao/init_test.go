package dao

import (
	"os"
	"testing"

	"github.com/teambition/urbs-console/src/service"
)

var (
	testDaos *Daos
)

func TestMain(m *testing.M) {
	testDaos = NewDaos(service.NewDB())
	os.Exit(m.Run())
}
