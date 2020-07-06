package http

import (
	"go-admin/internal/service"
	"net/http"
	"strings"
	"time"

	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/log"
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	pb "go-admin/api"
)

var svc pb.AdminServer

// New new a bm server.
func New(s pb.AdminServer) (engine *bm.Engine, err error) {
	var (
		cfg bm.ServerConfig
		ct  paladin.TOML
	)
	if err = paladin.Get("http.toml").Unmarshal(&ct); err != nil {
		return
	}
	if err = ct.Get("Server").UnmarshalTOML(&cfg); err != nil {
		return
	}
	svc = s
	ssvc := svc.(*service.Service)
	engine = bm.DefaultServer(&cfg)
	engine.Inject("/*", ssvc.SetToken())
	allowOriginHosts := []string{"localhost:9527"}
	config := &CORSConfig{
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           time.Duration(0),
		AllowOriginFunc: func(origin string) bool {
			for _, host := range allowOriginHosts {
				if strings.HasSuffix(strings.ToLower(origin), host) {
					return true
				}
			}
			return false
		},
	}
	engine.Use(newCORS(config))
	pb.RegisterAdminBMServer(engine, s)
	initRouter(engine)
	err = engine.Start()
	return
}

func initRouter(e *bm.Engine) {
	e.Ping(ping)
}

func ping(ctx *bm.Context) {
	if _, err := svc.Ping(ctx, nil); err != nil {
		log.Error("ping error(%v)", err)
		ctx.AbortWithStatus(http.StatusServiceUnavailable)
	}
}
