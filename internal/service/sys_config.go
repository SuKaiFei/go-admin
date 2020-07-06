package service

import (
	"context"
	"github.com/pkg/errors"
	pb "go-admin/api"
	"go-admin/internal/model"
)

// 通过字典类型获取字典数据
func (s *Service) GetConfigByKey(ctx context.Context, req *pb.ConfigByKeyReq) (resp *pb.ConfigInfoResp, err error) {
	config, err := s.dao.GetConfigByKey(ctx, req.Key)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	resp = convertConfigModelToPb(config)
	return
}

func convertConfigModelToPb(m *model.SysConfig) *pb.ConfigInfoResp {
	return &pb.ConfigInfoResp{
		Id:              int32(m.Id),
		Name:            m.Name,
		Key:             m.Key,
		Value:           m.Value,
		Type:            m.Type,
		Remark:          m.Remark,
		CreateBy:        m.CreateBy,
		UpdateBy:        m.UpdateBy,
		CreateTimestamp: m.CreateTimestamp,
		UpdateTimestamp: m.UpdateTimestamp,
		DeleteTimestamp: m.DeleteTimestamp,
	}
}
