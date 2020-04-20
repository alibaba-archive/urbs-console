package schema

import "time"

// User ...
type User struct {
	ID        int64     `gorm:"column:id" json:"-"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UID       string    `gorm:"column:uid" json:"uid"`             // varchar(63)，用户外部ID，表内唯一， 如 Teambition user id
	ActiveAt  int64     `gorm:"column:active_at" json:"active_at"` // 最近活跃时间戳，1970 以来的秒数，但不及时更新
	Labels    string    `gorm:"column:labels" json:"labels"`       // varchar(8190)，缓存用户当前被设置的 labels
}
