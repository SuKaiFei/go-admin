package service

import (
	"context"
	"github.com/pkg/errors"
	pb "go-admin/api"
	"go-admin/internal/model"
)

// 通过字典类型获取字典数据
func (s *Service) GetDictDataList(ctx context.Context, req *pb.DictDataListReq) (resp *pb.DictDataListResp, err error) {
	ds, err := s.dao.GetDictDataList(ctx, req)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	infos := make([]*pb.DictDataListResp_Info, len(ds))
	for i, d := range ds {
		infos[i] = convertDictDataModelToPb(d)
	}
	resp = &pb.DictDataListResp{Content: infos}
	return
}

func convertDictDataModelToPb(m *model.SysDictData) *pb.DictDataListResp_Info {
	return &pb.DictDataListResp_Info{
		Code:      int32(m.Code),
		Sort:      int32(m.Sort),
		Label:     m.Label,
		Value:     m.Value,
		Type:      m.Type,
		CssClass:  m.CssClass,
		ListClass: m.ListClass,
		IsDefault: m.IsDefault,
		State:     int32(m.State),
		Default:   m.Default,
		CreateBy:  m.CreateBy,
		UpdateBy:  m.UpdateBy,
		Remark:    m.Remark,
	}
}
