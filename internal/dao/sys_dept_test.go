package dao

import (
	"context"
	"github.com/go-kratos/kratos/pkg/log"
	"go-admin/internal/model"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDaoGetDeptList(t *testing.T) {
	Convey("GetDeptList", t, func() {
		var (
			ctx = context.Background()
			e   = &model.SysDept{}
		)
		Convey("When everything goes positive", func() {
			ds, err := d.GetDeptList(ctx, e)
			log.Info("ds(%v) err(%v)", ds, err)
			Convey("Then err should be nil.ds should not be nil.", func() {
				So(err, ShouldBeNil)
				So(ds, ShouldNotBeNil)
			})
		})
	})
}
