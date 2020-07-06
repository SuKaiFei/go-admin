package dao

import (
	"context"
	"github.com/go-kratos/kratos/pkg/cache/redis"
	"github.com/go-kratos/kratos/pkg/database/sql"
	"github.com/google/wire"
	pb "go-admin/api"
	"go-admin/internal/model"
)

var Provider = wire.NewSet(New, NewDB, NewRedis)

// Dao dao interface
type Dao interface {
	Close()
	Ping(ctx context.Context) (err error)

	// 根据用户名查找用户
	GetSysUserByUsername(ctx context.Context, username string) (user *model.SysUser, err error)
	// 根据用户名查找用户
	GetSysUser(ctx context.Context, userId int) (user *model.SysUser, err error)
	// 查询系统角色
	GetSysRole(ctx context.Context, roleId int) (role *model.SysRole, err error)
	// 查询系统角色列表
	GetSysRoles(ctx context.Context) (roles []*model.SysRole, err error)
	// 根据角色ID查询出有权限的菜单
	GetSysMenusByRoleId(ctx context.Context, roleId int) (menu []*model.SysMenu, err error)
	// 根据角色名称查询出有权限的菜单
	GetSysMenusByRoleName(ctx context.Context, roleName string) (menu []*model.SysMenu, err error)
	// 查询系统岗位列表
	GetSysPosts(ctx context.Context) (roles []*model.SysPost, err error)
	// 查询用户列表
	GetSysUsers(ctx context.Context, req *pb.UserListReq) (users []*model.SysUser, total int64, err error)
	// 查询部门列表
	GetDeptList(ctx context.Context, e *model.SysDept) (ds []*model.SysDept, err error)
	// 查询字典数据列表
	GetDictDataList(ctx context.Context, req *pb.DictDataListReq) (ds []*model.SysDictData, err error)
	// 根据key查询配置
	GetConfigByKey(ctx context.Context, key string) (config *model.SysConfig, err error)
	// 更新用户状态
	UpdateUserState(ctx context.Context, state pb.State, userId int) (updateCount int64, err error)
}

// dao dao.
type dao struct {
	db    *sql.DB
	redis *redis.Redis
}

// New new a dao and return.
func New(r *redis.Redis, db *sql.DB) (d Dao, cf func(), err error) {
	return newDao(r, db)
}

func newDao(r *redis.Redis, db *sql.DB) (d *dao, cf func(), err error) {
	d = &dao{
		db:    db,
		redis: r,
	}
	cf = d.Close
	return
}

// Close close the resource.
func (d *dao) Close() {
	d.db.Close()
}

// Ping ping the resource.
func (d *dao) Ping(ctx context.Context) (err error) {
	return nil
}
