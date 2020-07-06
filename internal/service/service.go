package service

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/mojocn/base64Captcha"
	"github.com/pkg/errors"
	pb "go-admin/api"
	"go-admin/internal/dao"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/wire"
)

var Provider = wire.NewSet(New, wire.Bind(new(pb.AdminServer), new(*Service)))

// Service service.
type Service struct {
	ac           *paladin.Map
	dao          dao.Dao
	AESBlock     cipher.Block
	jwtKey       []byte
	captchaStore base64Captcha.Store
}

// AESEncode aes encode
type AESEncode struct {
	AesKey string
	Salt   string
}

// New new a service and return.
func New(d dao.Dao) (s *Service, cf func(), err error) {
	s = &Service{
		ac:           &paladin.TOML{},
		dao:          d,
		captchaStore: base64Captcha.DefaultMemStore,
	}
	cf = s.Close
	if err = paladin.Watch("application.toml", s.ac); err != nil {
		return nil, nil, errors.WithStack(err)
	}
	aesE := new(AESEncode)
	if err = s.ac.Get("AESEncode").UnmarshalTOML(aesE); err != nil {
		return nil, nil, errors.WithStack(err)
	}
	s.AESBlock, err = aes.NewCipher([]byte(aesE.AesKey))
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	jwtKey, err := s.ac.Get("JwtKey").String()
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	s.jwtKey = []byte(jwtKey)
	return
}

// Ping ping the resource.
func (s *Service) Ping(ctx context.Context, e *empty.Empty) (*empty.Empty, error) {
	return &empty.Empty{}, s.dao.Ping(ctx)
}

// Close close the resource.
func (s *Service) Close() {
}
