package service

import (
	"context"
	pb "go-admin/api"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestServiceGetConfigByKey(t *testing.T) {
	Convey("GetConfigByKey", t, func() {
		var (
			ctx = context.Background()
			req = &pb.ConfigByKeyReq{
				Key: "sys.user.initPassword",
			}
		)
		Convey("When everything goes positive", func() {
			resp, err := s.GetConfigByKey(ctx, req)
			Convey("Then err should be nil.resp should not be nil.", func() {
				So(err, ShouldBeNil)
				So(resp, ShouldNotBeNil)
			})
		})
	})
}
