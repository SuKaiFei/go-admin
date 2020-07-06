package dao

import (
	"context"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDaoGetSysRole(t *testing.T) {
	Convey("GetSysRole", t, func() {
		var (
			ctx    = context.Background()
			roleId = int(0)
		)
		Convey("When everything goes positive", func() {
			role, err := d.GetSysRole(ctx, roleId)
			Convey("Then err should be nil.role should not be nil.", func() {
				So(err, ShouldBeNil)
				So(role, ShouldNotBeNil)
			})
		})
	})
}
