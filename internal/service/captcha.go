package service

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"

	pb "go-admin/api"

	"github.com/go-kratos/kratos/pkg/ecode"
	"github.com/google/uuid"
	"github.com/mojocn/base64Captcha"
)

// 获取验证码
func (s *Service) GetCaptcha(ctx context.Context, _ *empty.Empty) (resp *pb.CaptchaResp, err error) {
	id, b64s, err := s.driverDigit()
	if err != nil {
		return nil, ecode.Error(ecode.RequestErr, "验证码获取失败")
	}
	resp = &pb.CaptchaResp{
		Id:    id,
		Image: b64s,
	}
	return
}

//configJsonBody json request body.
type configJsonBody struct {
	Id            string
	CaptchaType   string
	VerifyValue   string
	DriverAudio   *base64Captcha.DriverAudio
	DriverString  *base64Captcha.DriverString
	DriverChinese *base64Captcha.DriverChinese
	DriverMath    *base64Captcha.DriverMath
	DriverDigit   *base64Captcha.DriverDigit
}

func (s *Service) driverDigit() (id, b64s string, err error) {
	e := configJsonBody{}
	e.Id = uuid.New().String()
	e.DriverDigit = base64Captcha.DefaultDriverDigit
	driver := e.DriverDigit
	return base64Captcha.NewCaptcha(driver, s.captchaStore).Generate()
}
