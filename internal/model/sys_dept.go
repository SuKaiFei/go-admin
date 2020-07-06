package model

type SysDept struct {
	Id       int    `ddb:"id" json:"id"`
	ParentId int    `ddb:"parent_id" json:"parent_id"` // 上级部门
	Path     string `ddb:"path" json:"path"`
	Name     string `ddb:"name" json:"name"`     // 部门名称
	Sort     int    `ddb:"sort" json:"sort"`     // 排序
	Leader   string `ddb:"leader" json:"leader"` // 负责人
	Phone    string `ddb:"phone" json:"phone"`   // 手机
	Email    string `ddb:"email" json:"email"`   // 邮箱
	State    int    `ddb:"state" json:"state"`   // 状态
	CreateBy string `ddb:"create_by" json:"create_by"`
	UpdateBy string `ddb:"update_by" json:"update_by"`

	CreateTimestamp int64 `json:"create_timestamp" ddb:"create_timestamp"`
	UpdateTimestamp int64 `json:"update_timestamp" ddb:"update_timestamp"`
	DeleteTimestamp int64 `json:"delete_timestamp" ddb:"delete_timestamp"`
}
