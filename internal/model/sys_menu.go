package model

type SysMenu struct {
	Id         int    `ddb:"id" json:"id"`
	Name       string `ddb:"name" json:"name"`
	Title      string `ddb:"title" json:"title"`
	Icon       string `ddb:"icon" json:"icon"`
	Path       string `ddb:"path" json:"path"`
	Paths      string `ddb:"paths" json:"paths"`
	Type       string `ddb:"type" json:"type"`
	Action     string `ddb:"action" json:"action"`
	Permission string `ddb:"permission" json:"permission"`
	ParentId   int    `ddb:"parent_id" json:"parent_id"`
	NoCache    bool   `ddb:"no_cache" json:"no_cache"`
	Breadcrumb string `ddb:"breadcrumb" json:"breadcrumb"`
	Component  string `ddb:"component" json:"component"`
	Sort       int    `ddb:"sort" json:"sort"`
	Visible    string `ddb:"visible" json:"visible"`
	CreateBy   string `ddb:"create_by" json:"create_by"`
	UpdateBy   string `ddb:"update_by" json:"update_by"`
	IsFrame    int    `ddb:"is_frame" json:"is_frame"`

	CreateTimestamp int64 `json:"create_timestamp" ddb:"create_timestamp"`
	UpdateTimestamp int64 `json:"update_timestamp" ddb:"update_timestamp"`
	DeleteTimestamp int64 `json:"delete_timestamp" ddb:"delete_timestamp"`
}
