package urbssetting

import "time"

// ModuleRes ...
type ModuleRes struct {
	SuccessResponseType
	Result *Module `json:"result"` // 空数组也保留
}

// ModulesRes ...
type ModulesRes struct {
	SuccessResponseType
	Result []*Module `json:"result"` // 空数组也保留
}

// Module ...
type Module struct {
	ID        int64      `gorm:"column:id" json:"-"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at"`
	OfflineAt *time.Time `gorm:"column:offline_at" json:"offline_at"` // 计划下线时间，用于灰度管理
	ProductID int64      `gorm:"column:product_id"`                   // 所从属的产品线 ID
	Name      string     `gorm:"column:name" json:"name"`             // varchar(63) 功能模块名称，产品线内唯一
	Desc      string     `gorm:"column:description" json:"desc"`      // varchar(1022) 功能模块描述
	Status    int64      `gorm:"column:status" json:"status"`         // -1 下线弃用，0 未使用，大于 0 为有效配置项数
}
