package service

import (
	"context"
	"github.com/bilibili/kratos/pkg/ecode"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	pb "go-admin/api"
	"go-admin/internal/model"
	"math"
)

// 获取用户列表数据
func (s *Service) GetUserList(ctx context.Context, req *pb.UserListReq) (resp *pb.UserListResp, err error) {
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}
	users, total, err := s.dao.GetSysUsers(ctx, req)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	resp = &pb.UserListResp{
		Total:     total,
		TotalPage: int64(math.Ceil(float64(total) / float64(req.PageSize))),
		Content:   convertUsersModelToPb(users),
	}
	return
}

// 更新用户状态
func (s *Service) UpdateUserState(ctx context.Context, req *pb.UpdateUserStateReq) (resp *empty.Empty, err error) {
	updateCount, err := s.dao.UpdateUserState(ctx, req.State, int(req.UserId))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if updateCount == 0 {
		return nil, ecode.Errorf(ecode.NothingFound, "找不到用户ID(%d)", req.UserId)
	}
	return
}

// 获取用户
func (s *Service) GetUser(ctx context.Context, req *pb.UserReq) (resp *pb.UserInfo, err error) {
	user, err := s.dao.GetSysUser(ctx, int(req.Id))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	resp = convertUserInfoModelToPb(user)
	return
}

func convertUserInfoModelToPb(m *model.SysUser) *pb.UserInfo {
	return &pb.UserInfo{
		Id:              int32(m.Id),
		NickName:        m.NickName,
		Phone:           m.Phone,
		RoleId:          int32(m.RoleId),
		Avatar:          m.Avatar,
		Sex:             m.Sex,
		Email:           m.Email,
		DeptId:          int32(m.DeptId),
		DeptName:        m.DeptName,
		PostId:          int32(m.PostId),
		CreateBy:        m.CreateBy,
		UpdateBy:        m.UpdateBy,
		Remark:          m.Remark,
		State:           int32(m.State),
		Username:        m.Username,
		CreateTimestamp: m.CreateTimestamp,
		UpdateTimestamp: m.UpdateTimestamp,
		DeleteTimestamp: m.DeleteTimestamp,
	}
}

func convertUsersModelToPb(ms []*model.SysUser) (ps []*pb.UserInfo) {
	ps = make([]*pb.UserInfo, len(ms))
	for i, m := range ms {
		ps[i] = convertUserInfoModelToPb(m)
	}
	return
}
