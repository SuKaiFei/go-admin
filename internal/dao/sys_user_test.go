package dao

import (
	"context"
	pb "go-admin/api"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDaoGetSysUserByUsername(t *testing.T) {
	Convey("GetSysUserByUsername", t, func() {
		var (
			ctx      = context.Background()
			username = ""
		)
		Convey("When everything goes positive", func() {
			user, err := d.GetSysUserByUsername(ctx, username)
			Convey("Then err should be nil.user should not be nil.", func() {
				So(err, ShouldBeNil)
				So(user, ShouldNotBeNil)
			})
		})
	})
}

func TestDaoGetSysUsers(t *testing.T) {
	Convey("GetSysUsers", t, func() {
		var (
			ctx = context.Background()
			req = &pb.UserListReq{}
		)
		Convey("When everything goes positive", func() {
			users, total, err := d.GetSysUsers(ctx, req)
			Convey("Then err should be nil.users,total should not be nil.", func() {
				So(err, ShouldBeNil)
				So(total, ShouldNotBeNil)
				So(users, ShouldNotBeNil)
			})
		})
	})
}

func TestDaoGetSysUser(t *testing.T) {
	Convey("GetSysUser", t, func() {
		var (
			ctx    = context.Background()
			userId = int(0)
		)
		Convey("When everything goes positive", func() {
			user, err := d.GetSysUser(ctx, userId)
			Convey("Then err should be nil.user should not be nil.", func() {
				So(err, ShouldBeNil)
				So(user, ShouldNotBeNil)
			})
		})
	})
}
