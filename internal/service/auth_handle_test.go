package service

import (
	"context"
	"go-admin/internal/model"
	"testing"

	"github.com/dgrijalva/jwt-go"
	. "github.com/smartystreets/goconvey/convey"
)

func TestServicesetJwtClaims(t *testing.T) {
	Convey("setJwtClaims", t, func() {
		var (
			ctx    = context.Background()
			claims jwt.MapClaims
			u      = &model.SysUser{}
		)
		Convey("When everything goes positive", func() {
			err := s.setJwtClaims(ctx, claims, u)
			Convey("Then err should be nil.", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}
