package service

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestServiceGetDeptTree(t *testing.T) {
	Convey("GetDeptTree", t, func() {
		var (
			ctx = _testAdminContext
		)
		Convey("When everything goes positive", func() {
			resp, err := s.GetDeptTree(ctx, nil)
			Convey("Then err should be nil.resp should not be nil.", func() {
				So(err, ShouldBeNil)
				So(resp, ShouldNotBeNil)
			})
		})
	})
}
