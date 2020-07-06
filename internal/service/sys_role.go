package service

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	pb "go-admin/api"
	"go-admin/internal/model"
)

// 获取角色列表
func (s *Service) GetRoleList(ctx context.Context, _ *empty.Empty) (resp *pb.RoleListResp, err error) {
	roleList, err := s.dao.GetSysRoles(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	infos := make([]*pb.RoleInfo, len(roleList))
	for i, role := range roleList {
		infos[i] = convertRoleModelToPb(role)
	}
	resp = &pb.RoleListResp{Content: infos}
	return
}

func convertRoleModelToPb(m *model.SysRole) *pb.RoleInfo {
	return &pb.RoleInfo{
		Id:              int32(m.Id),
		Name:            m.Name,
		State:           pb.State(m.State),
		Key:             m.Key,
		Sort:            int32(m.Sort),
		Flag:            m.Flag,
		CreateBy:        m.CreateBy,
		UpdateBy:        m.UpdateBy,
		Remark:          m.Remark,
		DataScope:       m.DataScope,
		Params:          m.Params,
		CreateTimestamp: m.CreateTimestamp,
		UpdateTimestamp: m.UpdateTimestamp,
		DeleteTimestamp: m.DeleteTimestamp,
	}
}
