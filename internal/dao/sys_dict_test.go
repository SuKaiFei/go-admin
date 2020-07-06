package dao

import (
	"context"
	pb "go-admin/api"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDaoGetDictDataList(t *testing.T) {
	Convey("GetDictDataList", t, func() {
		var (
			ctx = context.Background()
			req = &pb.DictDataListReq{
				DictType: "sys_normal_disable",
			}
		)
		Convey("When everything goes positive", func() {
			ds, err := d.GetDictDataList(ctx, req)
			Convey("Then err should be nil.ds should not be nil.", func() {
				So(err, ShouldBeNil)
				So(ds, ShouldNotBeNil)
			})
		})
	})
}
