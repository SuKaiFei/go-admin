// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package service

import (
	"github.com/google/wire"
	"go-admin/internal/dao"
)

//go:generate kratos tool wire
func newTestService() (*Service, func(), error) {
	panic(wire.Build(New, dao.Provider))
}
