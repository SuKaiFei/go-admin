package service

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	pb "go-admin/api"
	"go-admin/internal/model"
)

// 获取部门树
func (s *Service) GetDeptTree(ctx context.Context, _ *empty.Empty) (resp *pb.DeptTreeResp, err error) {
	deptList, err := s.dao.GetDeptList(ctx, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	resp = new(pb.DeptTreeResp)

	for _, dept := range deptList {
		if dept.ParentId != 0 {
			continue
		}
		resp.Content = append(resp.Content, RecursiveDept(&deptList, convertDeptModelToPb(dept)))
	}
	return
}

func RecursiveDept(deptList *[]*model.SysDept, dept *pb.DeptInfo) *pb.DeptInfo {
	list := *deptList

	min := make([]*pb.DeptInfo, 0)
	for _, deptM := range list {
		if dept.DeptId != int32(deptM.ParentId) {
			continue
		}
		ms := RecursiveDept(deptList, convertDeptModelToPb(deptM))
		min = append(min, ms)
	}
	dept.Children = min
	return dept
}

func convertDeptModelToPb(deptM *model.SysDept) *pb.DeptInfo {
	return &pb.DeptInfo{
		DeptId:   int32(deptM.Id),
		ParentId: int32(deptM.ParentId),
		DeptPath: deptM.Path,
		DeptName: deptM.Name,
		Sort:     int32(deptM.Sort),
		Leader:   deptM.Leader,
		Phone:    deptM.Phone,
		Email:    deptM.Email,
		State:    int32(deptM.State),
		Children: []*pb.DeptInfo{},
	}
}
