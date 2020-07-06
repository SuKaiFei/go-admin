package model

type SysPost struct {
	Id       int    `ddb:"id" json:"id"`
	Name     string `ddb:"name" json:"name"`
	Code     string `ddb:"code" json:"code"`
	Sort     int    `ddb:"sort" json:"sort"`
	State    int    `ddb:"state" json:"state"`
	Remark   string `ddb:"remark" json:"remark"`
	CreateBy string `ddb:"create_by" json:"create_by"`
	UpdateBy string `ddb:"update_by" json:"update_by"`

	CreateTimestamp int64 `ddb:"create_timestamp" json:"create_timestamp"`
	UpdateTimestamp int64 `ddb:"update_timestamp" json:"update_timestamp"`
	DeleteTimestamp int64 `ddb:"delete_timestamp" json:"delete_timestamp"`
}
