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
	_selDictDataList = "select * from sys_dict_data where state = 0 %s"
)

func (d *dao) GetDictDataList(ctx context.Context, req *pb.DictDataListReq) (ds []*model.SysDictData, err error) {
	var (
		args []interface{}
		cnd  string
	)
	if req != nil {
		if req.Code != "" {
			cnd += " and code = ? "
			args = append(args, req.Code)
		}
		if req.Label != "" {
			cnd += " and label = ? "
			args = append(args, req.Label)
		}
		if req.Type != "" {
			cnd += " and type = ? "
			args = append(args, req.Type)
		}
	}
	rows, err := d.db.Query(ctx, fmt.Sprintf(_selDictDataList, cnd), args...)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer rows.Close()
	if err = scanner.Scan(rows, &ds); err != nil {
		return nil, errors.WithStack(err)
	}
	return
}
