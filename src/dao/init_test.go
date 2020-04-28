package dao

import (
	"os"
	"testing"

	"github.com/teambition/urbs-console/src/service"
)

func TestMain(m *testing.M) {
	daos = NewDaos(service.NewDB())
	os.Exit(m.Run())
}
