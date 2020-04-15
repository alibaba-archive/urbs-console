package dao

import (
	"github.com/teambition/urbs-console/src/service"
	"github.com/teambition/urbs-console/src/util"
)

func init() {
	util.DigProvide(NewDaos)
}

// Daos ...
type Daos struct {
	OperationLog *OperationLog
}

// NewDaos ...
func NewDaos(sql *service.SQL) *Daos {
	return &Daos{
		OperationLog: &OperationLog{DB: sql.DB},
	}
}
