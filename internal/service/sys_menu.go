package service

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	pb "go-admin/api"
	"go-admin/internal/model"
)

// 根据角色名称获取菜单列表数据（左菜单使用）
func (s *Service) GetMenuByRoleKey(ctx context.Context, _ *empty.Empty) (resp *pb.MenuRoleResp, err error) {
	roleName, err := s.getRoleKey(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	menus, err := s.dao.GetSysMenusByRoleName(ctx, roleName)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	resp = new(pb.MenuRoleResp)

	resp.Content = make([]*pb.MenuRoleResp_Info, 0)
	for i := 0; i < len(menus); i++ {
		if menus[i].ParentId != 0 {
			continue
		}
		menusInfo := RecursiveMenu(&menus, convertMenuModelToPb(menus[i]))
		resp.Content = append(resp.Content, menusInfo)
	}
	return
}

func RecursiveMenu(menus *[]*model.SysMenu, menu *pb.MenuRoleResp_Info) *pb.MenuRoleResp_Info {
	list := *menus
	min := make([]*pb.MenuRoleResp_Info, 0)

	for j := 0; j < len(list); j++ {
		if menu.MenuId != int32(list[j].ParentId) {
			continue
		}
		mi := &pb.MenuRoleResp_Info{
			MenuId:     int32(list[j].Id),
			MenuName:   list[j].Name,
			Title:      list[j].Title,
			Icon:       list[j].Icon,
			Path:       list[j].Path,
			MenuType:   list[j].Type,
			Action:     list[j].Action,
			Permission: list[j].Permission,
			ParentId:   int32(list[j].ParentId),
			NoCache:    list[j].NoCache,
			Breadcrumb: list[j].Breadcrumb,
			Component:  list[j].Component,
			Sort:       int32(list[j].Sort),
			Visible:    list[j].Visible,
			Children:   []*pb.MenuRoleResp_Info{},
		}
		if mi.MenuType != "F" {
			ms := RecursiveMenu(menus, mi)
			min = append(min, ms)
		} else {
			min = append(min, mi)
		}
	}
	menu.Children = min
	return menu
}

func convertMenuModelToPb(m *model.SysMenu) *pb.MenuRoleResp_Info {
	return &pb.MenuRoleResp_Info{
		MenuId:     int32(m.Id),
		MenuName:   m.Name,
		Title:      m.Title,
		Icon:       m.Icon,
		Path:       m.Path,
		MenuType:   m.Type,
		Action:     m.Action,
		Permission: m.Permission,
		ParentId:   int32(m.ParentId),
		NoCache:    m.NoCache,
		Breadcrumb: m.Breadcrumb,
		Component:  m.Component,
		Sort:       int32(m.Sort),
		Visible:    m.Visible,
		Children:   []*pb.MenuRoleResp_Info{},
	}
}
