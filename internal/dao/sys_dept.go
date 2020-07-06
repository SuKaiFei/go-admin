package dao

import (
	"context"
	"fmt"
	"github.com/didi/gendry/scanner"
	"github.com/pkg/errors"
	"go-admin/internal/model"
)

const (
	_selDeptList = "select * from sys_dept where 1=1 %s"
)

func (d *dao) GetDeptList(ctx context.Context, e *model.SysDept) (ds []*model.SysDept, err error) {
	var (
		args []interface{}
		cnd  string
	)
	if e != nil {
		if e.Id != 0 {
			cnd += " and id = ?"
			args = append(args, e.Id)
		}
		if e.Name != "" {
			cnd += " and name = ?"
			args = append(args, e.Name)
		}
		if e.Path != "" {
			cnd += " and path = ?"
			args = append(args, e.Path)
		}
	}
	rows, err := d.db.Query(ctx, fmt.Sprintf(_selDeptList, cnd), args...)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer rows.Close()
	if err = scanner.Scan(rows, &ds); err != nil {
		return nil, errors.WithStack(err)
	}
	return
}
