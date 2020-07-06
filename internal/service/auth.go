package service

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/didi/gendry/scanner"
	"github.com/go-kratos/kratos/pkg/ecode"
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	pb "go-admin/api"
	"golang.org/x/crypto/bcrypt"
)

const (
	_defaultAvatar = "http://img.crcz.com/allimg/202001/23/1579793244834816.jpg"
)

// 登录
func (s *Service) Login(ctx context.Context, req *pb.LoginReq) (resp *pb.LoginResp, err error) {
	if !s.captchaStore.Verify(req.Uuid, req.Code, true) {
		return nil, ecode.Error(ecode.RequestErr, "验证码错误")
	}
	user, err := s.dao.GetSysUserByUsername(ctx, req.Username)
	if err != nil {
		if errors.Cause(err) == scanner.ErrEmptyResult {
			return nil, ecode.Error(ecode.RequestErr, "用户不存在")
		}
		return nil, errors.WithStack(err)
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, ecode.Error(ecode.RequestErr, "用户名或者密码错误")
	}

	token := jwt.New(jwt.GetSigningMethod("HS256"))
	claims := token.Claims.(jwt.MapClaims)
	if err = s.setJwtClaims(ctx, claims, user); err != nil {
		return nil, errors.WithStack(err)
	}
	tokenStr, err := token.SignedString(s.jwtKey)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	resp = &pb.LoginResp{
		Expire: claims[_jwtClaimsExp].(int64),
		Token:  tokenStr,
	}
	return
}

// 登出
func (s *Service) Logout(ctx context.Context, _ *empty.Empty) (resp *empty.Empty, err error) {
	return &empty.Empty{}, nil
}

// 获取用户信息
func (s *Service) GetUserInfo(ctx context.Context, _ *empty.Empty) (resp *pb.UserInfoResp, err error) {
	bmContext := ctx.(*bm.Context)
	tokenI, _ := bmContext.Get("JWT_TOKEN")
	token := tokenI.(*jwt.Token)
	user, permissions, roles, err := s.getJwtClaims(ctx, token.Claims.(jwt.MapClaims))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if user.Avatar == "" {
		user.Avatar = _defaultAvatar
	}
	resp = &pb.UserInfoResp{
		Avatar:       user.Avatar,
		Introduction: user.Remark,
		Name:         user.NickName,
		Permissions:  permissions,
		Roles:        roles,
	}
	return
}
