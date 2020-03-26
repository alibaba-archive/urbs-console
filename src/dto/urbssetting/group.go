package urbssetting

import "time"

// GroupsBody ...
type GroupsBody struct {
	Groups []*GroupBody `json:"groups"`
}

// GroupBody ...
type GroupBody struct {
	UID  string `json:"uid"`
	Kind string `json:"kind"`
	Desc string `json:"desc"`
}

// GroupsURL ...
type GroupsURL struct {
	Pagination
	Kind string `json:"kind" query:"kind"`
}

// GroupsRes ...
type GroupsRes struct {
	SuccessResponseType
	Result []*Group `json:"result"`
}

// Group ...
type Group struct {
	ID        int64     `gorm:"column:id" json:"-"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	SyncAt    int64     `gorm:"column:sync_at" json:"sync_at"`  // 群组成员同步时间点，1970 以来的秒数
	UID       string    `gorm:"column:uid" json:"uid"`          // varchar(63)，群组外部ID，表内唯一， 如 Teambition organization id
	Kind      string    `gorm:"column:kind" json:"kind"`        // varchar(63)，群组外部ID，表内唯一， 如 Teambition organization id
	Desc      string    `gorm:"column:description" json:"desc"` // varchar(1022)，群组描述
}

// GroupRes ...
type GroupRes struct {
	SuccessResponseType
	Result *Group `json:"result"`
}

// GroupMember ...
type GroupMember struct {
	ID        int64     `json:"-"`
	User      string    `json:"user"`
	CreatedAt time.Time `json:"created_at"`
	SyncAt    int64     `json:"sync_at"` // 归属关系同步时间戳，1970 以来的秒数，应该与 group.sync_at 相等
}

// GroupMembersRes ...
type GroupMembersRes struct {
	SuccessResponseType
	Result []*GroupMember `json:"result"`
}

// GroupMembersURL ...
type GroupMembersURL struct {
	UID    string `json:"uid" param:"uid"`
	User   string `json:"user" query:"user"`       // 根据用户 uid 删除一个成员
	SyncLt int64  `json:"sync_lt" query:"sync_lt"` // 或根据 sync_lt 删除同步时间小于指定值的所有成员
}

// GroupUpdateBody ...
type GroupUpdateBody struct {
	Desc   *string `json:"desc"`
	SyncAt *int64  `json:"sync_at"`
}
