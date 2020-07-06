// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"go-admin/internal/dao"
	"go-admin/internal/server/grpc"
	"go-admin/internal/server/http"
	"go-admin/internal/service"

	"github.com/google/wire"
)

//go:generate kratos t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, http.New, grpc.New, NewApp))
}
