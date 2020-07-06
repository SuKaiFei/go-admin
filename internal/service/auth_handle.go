package service

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-kratos/kratos/pkg/ecode"
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	"github.com/pkg/errors"
	pb "go-admin/api"
	"go-admin/internal/model"
	"strconv"
	"time"
)

const (
	// jwt map claims
	_jwtClaimsIdentityKey  = "identity"
	_jwtClaimsNiceKey      = "nice"
	_jwtClaimsDataScopeKey = "data_scope"
	_jwtClaimsRoleIdKey    = "role_id"
	_jwtClaimsRoleKey      = "role_key"
	_jwtClaimsRoleNameKey  = "role_name"
	_jwtClaimsExp          = "exp"
	_jwtClaimsOrigIat      = "orig_iat"

	_jwtTimeOut = time.Hour * 2
)

func (s *Service) setJwtClaims(ctx context.Context, claims jwt.MapClaims, u *model.SysUser) (err error) {
	uId, err := s.encrypt(strconv.Itoa(u.Id))
	if err != nil {
		return errors.WithStack(err)
	}
	r, err := s.dao.GetSysRole(ctx, u.RoleId)
	if err != nil {
		return errors.WithStack(err)
	}
	rId, err := s.encrypt(strconv.Itoa(u.RoleId))
	if err != nil {
		return errors.WithStack(err)
	}
	expire := time.Now().Add(_jwtTimeOut)
	claims[_jwtClaimsExp] = expire.Unix()
	claims[_jwtClaimsOrigIat] = time.Now().Unix()
	claims[_jwtClaimsIdentityKey] = uId
	claims[_jwtClaimsRoleIdKey] = rId
	claims[_jwtClaimsRoleKey] = r.Key
	claims[_jwtClaimsNiceKey] = u.Username
	claims[_jwtClaimsDataScopeKey] = r.DataScope
	claims[_jwtClaimsRoleNameKey] = r.Name
	return
}

func (s *Service) getJwtClaims(ctx context.Context, claims jwt.MapClaims) (u *model.SysUser, permissions []string, roles []string, err error) {
	uIdStr, err := s.decrypt(claims[_jwtClaimsIdentityKey].(string))
	if err != nil {
		return nil, nil, nil, errors.WithStack(err)
	}
	uId, err := strconv.Atoi(uIdStr)
	if err != nil {
		return nil, nil, nil, errors.WithStack(err)
	}
	u, err = s.dao.GetSysUser(ctx, uId)
	if err != nil {
		return nil, nil, nil, errors.WithStack(err)
	}
	roleName := claims[_jwtClaimsRoleKey]
	if roleName == "admin" || roleName == "系统管理员" {
		permissions = []string{"*:*:*"}
	} else {
		rIdStr, err := s.decrypt(claims[_jwtClaimsRoleIdKey].(string))
		if err != nil {
			return nil, nil, nil, errors.WithStack(err)
		}
		rId, err := strconv.Atoi(rIdStr)
		if err != nil {
			return nil, nil, nil, errors.WithStack(err)
		}
		menus, err := s.dao.GetSysMenusByRoleId(ctx, rId)
		if err != nil {
			return nil, nil, nil, errors.WithStack(err)
		}
		permissions = make([]string, len(menus))
		for i, menu := range menus {
			permissions[i] = menu.Permission
		}
	}
	roles = []string{roleName.(string)}
	return
}

func (s *Service) getRoleKey(ctx context.Context) (roleName string, err error) {
	bmContext := ctx.(*bm.Context)
	tokenI, _ := bmContext.Get("JWT_TOKEN")
	token := tokenI.(*jwt.Token)
	mapClaims := token.Claims.(jwt.MapClaims)
	roleNameI := mapClaims[_jwtClaimsRoleKey]
	roleName = roleNameI.(string)
	return
}

var (
	_whitelistSetToken = map[string]struct{}{
		pb.PathAdminLogin:      {},
		pb.PathAdminGetCaptcha: {},
	}
)

func (s *Service) SetToken() bm.HandlerFunc {
	return func(c *bm.Context) {
		if c.Request.Method == "OPTIONS" {
			c.Next()
			return
		}
		_, ok := _whitelistSetToken[c.RoutePath]
		if ok || len(c.Request.Header.Get("Authorization")) < 7 {
			c.Next()
			return
		}
		token, err := jwt.Parse(c.Request.Header.Get("Authorization")[7:], func(token *jwt.Token) (interface{}, error) {
			return s.jwtKey, nil
		})
		if err != nil {
			if errors.Cause(err).Error() == "Token is expired" {
				c.JSON(nil, ecode.Error(ecode.Unauthorized, err.Error()))
			} else {
				c.JSON(nil, ecode.Error(ecode.RequestErr, err.Error()))
			}
			c.Abort()
			return
		}
		if !token.Valid {
			c.JSON(nil, ecode.Error(ecode.RequestErr, "Token is not valid"))
			c.Abort()
			return
		}
		c.Set("JWT_TOKEN", token)
		c.Next()
	}
}
