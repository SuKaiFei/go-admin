package model

type SysRole struct {
	Id       int    `ddb:"id" json:"id"`
	Name     string `ddb:"name" json:"name"`
	State    int    `ddb:"state" json:"state"`
	Key      string `ddb:"key" json:"key"`   // 角色代码
	Sort     int    `ddb:"sort" json:"sort"` // 角色排序
	Flag     string `ddb:"flag" json:"flag"`
	CreateBy string `ddb:"create_by" json:"create_by"`
	UpdateBy string `ddb:"update_by" json:"update_by"`
	Remark   string `ddb:"remark" json:"remark"` // 备注

	DataScope string `json:"data_scope"`
	Params    string `json:"params"`

	CreateTimestamp int64 `json:"create_timestamp" ddb:"create_timestamp"`
	UpdateTimestamp int64 `json:"update_timestamp" ddb:"update_timestamp"`
	DeleteTimestamp int64 `json:"delete_timestamp" ddb:"delete_timestamp"`
}
