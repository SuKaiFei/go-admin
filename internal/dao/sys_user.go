package dao

import (
	"context"
	"fmt"
	"github.com/didi/gendry/scanner"
	"github.com/pkg/errors"
	pb "go-admin/api"
	"go-admin/internal/model"
)

const (
	_selUserByUsername = "select * from sys_user where username = ? limit 1"
	_selUsers          = "select sys_user.*,sys_dept.name dept_name from sys_user left join sys_dept on sys_dept.id = sys_user.dept_id where 1=1 %s limit ?,?"
	_countUsers        = "select count(*) from sys_user left join sys_dept on sys_dept.id = sys_user.dept_id where 1=1 %s "
	_updateUserState   = "update sys_user set `state`=?, `update_timestamp`=UNIX_TIMESTAMP() where id = ?"
	_selUserById       = "select * from sys_user where id = ? limit 1"
)

func (d *dao) GetSysUserByUsername(ctx context.Context, username string) (user *model.SysUser, err error) {
	rows, err := d.db.Query(ctx, _selUserByUsername, username)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer rows.Close()
	user = new(model.SysUser)
	if err = scanner.Scan(rows, user); err != nil {
		return nil, errors.WithStack(err)
	}
	return
}

func (d *dao) UpdateUserState(ctx context.Context, state pb.State, userId int) (updateCount int64, err error) {
	rows, err := d.db.Exec(ctx, _updateUserState, state, userId)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	if updateCount, err = rows.RowsAffected(); err != nil {
		return 0, errors.WithStack(err)
	}
	return
}

func (d *dao) GetSysUsers(ctx context.Context, req *pb.UserListReq) (users []*model.SysUser, total int64, err error) {
	var (
		cnd  = ""
		args []interface{}
	)
	if req.DeptId != 0 {
		cnd += " and sys_dept.id = ?"
		args = append(args, req.DeptId)
	}

	row := d.db.QueryRow(ctx, fmt.Sprintf(_countUsers, cnd), args...)
	if err := row.Scan(&total); err != nil {
		return nil, 0, errors.WithStack(err)
	}
	if total == 0 {
		return
	}
	args = append(args, (req.Page-1)*req.PageSize, req.PageSize)
	rows, err := d.db.Query(ctx, fmt.Sprintf(_selUsers, cnd), args...)
	if err != nil {
		return nil, 0, errors.WithStack(err)
	}
	defer rows.Close()
	if err = scanner.Scan(rows, &users); err != nil {
		return nil, 0, errors.WithStack(err)
	}
	return
}

func (d *dao) GetSysUser(ctx context.Context, userId int) (user *model.SysUser, err error) {
	rows, err := d.db.Query(ctx, _selUserById, userId)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer rows.Close()
	user = new(model.SysUser)
	if err = scanner.Scan(rows, user); err != nil {
		return nil, errors.WithStack(err)
	}
	return
}
