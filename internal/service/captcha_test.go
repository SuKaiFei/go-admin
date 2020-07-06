package service

import (
	"context"
	"testing"

	"github.com/golang/protobuf/ptypes/empty"
	. "github.com/smartystreets/goconvey/convey"
)

func TestServiceGetCaptcha(t *testing.T) {
	Convey("GetCaptcha", t, func() {
		var (
			ctx = context.Background()
			req = &empty.Empty{}
		)
		Convey("When everything goes positive", func() {
			resp, err := s.GetCaptcha(ctx, req)
			Convey("Then err should be nil.resp should not be nil.", func() {
				So(err, ShouldBeNil)
				So(resp, ShouldNotBeNil)
			})
		})
	})
}

func TestServicedriverDigit(t *testing.T) {
	Convey("driverDigit", t, func() {
		Convey("When everything goes positive", func() {
			id, b64s, err := s.driverDigit()
			Convey("Then err should be nil.id,b64s should not be nil.", func() {
				So(err, ShouldBeNil)
				So(b64s, ShouldNotBeNil)
				So(id, ShouldNotBeNil)
			})
		})
	})
}
