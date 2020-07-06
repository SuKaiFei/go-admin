package service

import (
	"github.com/golang/protobuf/ptypes/empty"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestServiceGetMenuByRoleKey(t *testing.T) {
	Convey("GetMenuByRoleKey", t, func() {
		var (
			ctx = _testAdminContext
			req = &empty.Empty{}
		)
		Convey("When everything goes positive", func() {
			resp, err := s.GetMenuByRoleKey(ctx, req)
			Convey("Then err should be nil.resp should not be nil.", func() {
				So(err, ShouldBeNil)
				So(resp, ShouldNotBeNil)
			})
		})
	})
}
