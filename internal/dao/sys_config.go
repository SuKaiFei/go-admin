package dao

import (
	"context"
	"github.com/didi/gendry/scanner"
	"github.com/pkg/errors"
	"go-admin/internal/model"
)

const (
	_selConfigByKey = "select * from sys_config where `key` = ? limit 1"
)

func (d *dao) GetConfigByKey(ctx context.Context, key string) (config *model.SysConfig, err error) {
	rows, err := d.db.Query(ctx, _selConfigByKey, key)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer rows.Close()
	config = new(model.SysConfig)
	if err = scanner.Scan(rows, config); err != nil {
		return nil, errors.WithStack(err)
	}
	return
}
