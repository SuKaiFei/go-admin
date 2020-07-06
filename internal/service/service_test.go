package service

import (
	"context"
	"flag"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	"os"
	"testing"
)

var (
	s *Service

	_testAdminContext = &bm.Context{}
	_testAdminToken   = &jwt.Token{Claims: jwt.MapClaims{
		_jwtClaimsRoleKey: "admin",
	}}
)

func TestMain(m *testing.M) {
	flag.Set("conf", "../../configs")
	flag.Parse()
	var err error
	if err = paladin.Init(); err != nil {
		panic(err)
	}
	var cf func()
	if s, cf, err = newTestService(); err != nil {
		panic(err)
	}
	_testAdminContext.Context = context.Background()
	_testAdminContext.Set("JWT_TOKEN", _testAdminToken)
	ret := m.Run()
	cf()
	os.Exit(ret)
}
