package service

import (
	"context"
	pb "go-admin/api"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestServiceLogin(t *testing.T) {
	Convey("Login", t, func() {
		var (
			ctx = context.Background()
			req = &pb.LoginReq{}
		)
		Convey("When everything goes positive", func() {
			resp, err := s.Login(ctx, req)
			Convey("Then err should be nil.resp should not be nil.", func() {
				So(err, ShouldBeNil)
				So(resp, ShouldNotBeNil)
			})
		})
	})
}
