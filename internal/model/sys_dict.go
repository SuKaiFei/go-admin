package model

type SysDictType struct {
	Id       int    `ddb:"id" json:"id"`
	Name     string `ddb:"name" json:"name"`   // 字典名称
	Type     string `ddb:"type" json:"type"`   // 字典类型
	State    int    `ddb:"state" json:"state"` // 状态
	CreateBy string `ddb:"create_by" json:"create_by"`
	UpdateBy string `ddb:"update_by" json:"update_by"`
	Remark   string `ddb:"remark" json:"remark"`

	CreateTimestamp int64 `json:"create_timestamp" ddb:"create_timestamp"`
	UpdateTimestamp int64 `json:"update_timestamp" ddb:"update_timestamp"`
	DeleteTimestamp int64 `json:"delete_timestamp" ddb:"delete_timestamp"`
}

type SysDictData struct {
	Code      int    `ddb:"code" json:"code"`   // 字典编码
	Sort      int    `ddb:"sort" json:"sort"`   // 显示顺序
	Label     string `ddb:"label" json:"label"` // 数据标签
	Value     string `ddb:"value" json:"value"` // 数据键值
	Type      string `ddb:"type" json:"type"`   // 字典类型
	CssClass  string `ddb:"css_class" json:"css_class"`
	ListClass string `ddb:"list_class" json:"list_class"`
	IsDefault string `ddb:"is_default" json:"is_default"`
	State     int    `ddb:"state" json:"state"` // 状态
	Default   string `ddb:"default" json:"default"`
	CreateBy  string `ddb:"create_by" json:"create_by"`
	UpdateBy  string `ddb:"update_by" json:"update_by"`
	Remark    string `ddb:"remark" json:"remark"` // 备注

	CreateTimestamp int64 `json:"create_timestamp" ddb:"create_timestamp"`
	UpdateTimestamp int64 `json:"update_timestamp" ddb:"update_timestamp"`
	DeleteTimestamp int64 `json:"delete_timestamp" ddb:"delete_timestamp"`
}
