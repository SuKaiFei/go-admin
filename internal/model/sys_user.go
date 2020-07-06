package model

type SysUser struct {
	Id       int    `ddb:"id" json:"id"`
	NickName string `ddb:"nick_name" json:"nick_name"` // 昵称
	Phone    string `ddb:"phone" json:"phone"`         // 手机号
	RoleId   int    `ddb:"role_id" json:"role_id"`     // 角色编码
	Avatar   string `ddb:"avatar" json:"avatar"`       // 头像
	Sex      string `ddb:"sex" json:"sex"`             // 性别
	Email    string `ddb:"email" json:"email"`         // 邮箱
	DeptId   int    `ddb:"dept_id" json:"dept_id"`     // 部门编码
	DeptName string `ddb:"dept_name" json:"dept_name"` // 部门名称
	PostId   int    `ddb:"post_id" json:"post_id"`     // 职位编码
	CreateBy string `ddb:"create_by" json:"create_by"`
	UpdateBy string `ddb:"update_by" json:"update_by"`
	Remark   string `ddb:"remark" json:"remark"` // 备注
	State    int    `ddb:"state" json:"state"`
	Username string `ddb:"username" json:"username"`
	Password string `ddb:"password" json:"password"`

	CreateTimestamp int64 `json:"create_timestamp" ddb:"create_timestamp"`
	UpdateTimestamp int64 `json:"update_timestamp" ddb:"update_timestamp"`
	DeleteTimestamp int64 `json:"delete_timestamp" ddb:"delete_timestamp"`
}
