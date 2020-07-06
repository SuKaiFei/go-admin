package service

import (
	pb "go-admin/api"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestServiceGetDictDataByType(t *testing.T) {
	Convey("GetDictDataByType", t, func() {
		var (
			ctx = _testAdminContext
			req = &pb.DictDataListReq{
				DictType: "sys_normal_disable",
			}
		)
		Convey("When everything goes positive", func() {
			resp, err := s.GetDictDataList(ctx, req)
			Convey("Then err should be nil.resp should not be nil.", func() {
				So(err, ShouldBeNil)
				So(resp, ShouldNotBeNil)
			})
		})
	})
}
