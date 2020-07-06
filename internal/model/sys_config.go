package model

type SysConfig struct {
	Id       int    `ddb:"id" json:"id"`
	Name     string `ddb:"name" json:"name"`     // 参数名称
	Key      string `ddb:"key" json:"key"`       // 参数键名
	Value    string `ddb:"value" json:"value"`   // 参数键值
	Type     string `ddb:"type" json:"type"`     // 是否系统内置
	Remark   string `ddb:"remark" json:"remark"` // 备注
	CreateBy string `ddb:"create_by" json:"create_by"`
	UpdateBy string `ddb:"update_by" json:"update_by"`

	CreateTimestamp int64 `json:"create_timestamp" ddb:"create_timestamp"`
	UpdateTimestamp int64 `json:"update_timestamp" ddb:"update_timestamp"`
	DeleteTimestamp int64 `json:"delete_timestamp" ddb:"delete_timestamp"`
}
