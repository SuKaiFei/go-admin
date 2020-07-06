package dao

import (
	"context"
	"github.com/didi/gendry/scanner"
	"github.com/pkg/errors"
	"go-admin/internal/model"
)

const (
	_selPermissionMenu  = "select * from sys_role_menu left join sys_menu on sys_menu.id = sys_role_menu.menu_id where sys_menu.type in('F','C') and role_id = ?"
	_selMenusByRoleName = "select * from sys_menu left join sys_role_menu on sys_role_menu.menu_id=sys_menu.id where sys_role_menu.role_name=? and  type in ('M','C') order by sort asc"
)

func (d *dao) GetSysMenusByRoleId(ctx context.Context, roleId int) (menu []*model.SysMenu, err error) {
	rows, err := d.db.Query(ctx, _selPermissionMenu, roleId)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer rows.Close()
	if err = scanner.Scan(rows, &menu); err != nil {
		return nil, errors.WithStack(err)
	}
	return
}

func (d *dao) GetSysMenusByRoleName(ctx context.Context, roleName string) (menu []*model.SysMenu, err error) {
	rows, err := d.db.Query(ctx, _selMenusByRoleName, roleName)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer rows.Close()
	if err = scanner.Scan(rows, &menu); err != nil {
		return nil, errors.WithStack(err)
	}
	return
}
