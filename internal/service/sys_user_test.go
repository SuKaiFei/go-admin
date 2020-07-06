package service

import (
	"context"
	pb "go-admin/api"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestServiceGetUserList(t *testing.T) {
	Convey("GetUserList", t, func() {
		var (
			ctx = context.Background()
			req = &pb.UserListReq{}
		)
		Convey("When everything goes positive", func() {
			resp, err := s.GetUserList(ctx, req)
			Convey("Then err should be nil.resp should not be nil.", func() {
				So(err, ShouldBeNil)
				So(resp, ShouldNotBeNil)
			})
		})
	})
}
