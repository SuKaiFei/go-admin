package service

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	pb "go-admin/api"
	"go-admin/internal/model"
)

// 获取角色列表
func (s *Service) GetPostList(ctx context.Context, _ *empty.Empty) (resp *pb.PostListResp, err error) {
	postList, err := s.dao.GetSysPosts(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	infos := make([]*pb.PostInfo, len(postList))
	for i, post := range postList {
		infos[i] = convertPostModelToPb(post)
	}
	resp = &pb.PostListResp{Content: infos}
	return
}

func convertPostModelToPb(m *model.SysPost) *pb.PostInfo {
	return &pb.PostInfo{
		Id:              int32(m.Id),
		Name:            m.Name,
		Code:            m.Code,
		Sort:            int32(m.Sort),
		State:           pb.State(m.State),
		Remark:          m.Remark,
		CreateBy:        m.CreateBy,
		UpdateBy:        m.UpdateBy,
		CreateTimestamp: m.CreateTimestamp,
		UpdateTimestamp: m.UpdateTimestamp,
		DeleteTimestamp: m.DeleteTimestamp,
	}
}
