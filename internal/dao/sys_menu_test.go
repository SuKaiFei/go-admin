package dao

import (
	"context"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDaoGetSysMenusByRoleId(t *testing.T) {
	Convey("GetSysMenusByRoleId", t, func() {
		var (
			ctx    = context.Background()
			roleId = int(0)
		)
		Convey("When everything goes positive", func() {
			menu, err := d.GetSysMenusByRoleId(ctx, roleId)
			Convey("Then err should be nil.menu should not be nil.", func() {
				So(err, ShouldBeNil)
				So(menu, ShouldNotBeNil)
			})
		})
	})
}

func TestDaoGetSysMenusByRoleName(t *testing.T) {
	Convey("GetSysMenusByRoleName", t, func() {
		var (
			ctx      = context.Background()
			roleName = ""
		)
		Convey("When everything goes positive", func() {
			menu, err := d.GetSysMenusByRoleName(ctx, roleName)
			Convey("Then err should be nil.menu should not be nil.", func() {
				So(err, ShouldBeNil)
				So(menu, ShouldNotBeNil)
			})
		})
	})
}
