package dao

import (
	"context"
	"github.com/didi/gendry/scanner"
	"github.com/pkg/errors"
	"go-admin/internal/model"
)

const (
	_selSysPosts = "select * from sys_post where state = 0 order by sort asc"
)

func (d *dao) GetSysPosts(ctx context.Context) (posts []*model.SysPost, err error) {
	rows, err := d.db.Query(ctx, _selSysPosts)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer rows.Close()
	if err = scanner.Scan(rows, &posts); err != nil {
		return nil, errors.WithStack(err)
	}
	return
}
