package dao

import (
	"context"
	"github.com/didi/gendry/scanner"
	"github.com/pkg/errors"
	"go-admin/internal/model"
)

const (
	_selSysRoleById = "select * from sys_role where id = ? limit 1"
	_selSysRoles    = "select * from sys_role where state = 0 order by sort asc"
)

func (d *dao) GetSysRole(ctx context.Context, roleId int) (role *model.SysRole, err error) {
	rows, err := d.db.Query(ctx, _selSysRoleById, roleId)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer rows.Close()
	role = new(model.SysRole)
	if err = scanner.Scan(rows, role); err != nil {
		return nil, errors.WithStack(err)
	}
	return
}

func (d *dao) GetSysRoles(ctx context.Context) (roles []*model.SysRole, err error) {
	rows, err := d.db.Query(ctx, _selSysRoles)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer rows.Close()
	if err = scanner.Scan(rows, &roles); err != nil {
		return nil, errors.WithStack(err)
	}
	return
}
