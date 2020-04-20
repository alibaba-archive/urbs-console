package schema

import "time"

// Group ...
type Group struct {
	ID        int64     `gorm:"column:id" json:"-"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	SyncAt    int64     `gorm:"column:sync_at" json:"sync_at"`  // 群组成员同步时间点，1970 以来的秒数
	UID       string    `gorm:"column:uid" json:"uid"`          // varchar(63)，群组外部ID，表内唯一， 如 Teambition organization id
	Kind      string    `gorm:"column:kind" json:"kind"`        // varchar(63)，群组外部ID，表内唯一， 如 Teambition organization id
	Desc      string    `gorm:"column:description" json:"desc"` // varchar(1022)，群组描述
	Status    int64     `gorm:"column:status" json:"status"`    // 成员计数，非实时精确值
}
