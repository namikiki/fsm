package di

import (
	"context"
	"crypto/md5"
	"fsm/middlewares"
	"fsm/pkg/configs"
	"fsm/pkg/database"
	"fsm/pkg/minio"
	"fsm/pkg/redis"
	"fsm/pkg/repositories"
	"fsm/router"
	"fsm/router/controllers"
	"fsm/services"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"hash"
	"log"
	"time"
)

var RepositoriesModule = fx.Module("RepositoriesModule",
	fx.Provide(
		repositories.NewStorageRepository,
		repositories.NewUserRepository,
	),
)

var PKGModule = fx.Module("pkg",
	fx.Provide(
		configs.NewConfig,
		redis.NewRedis,
		minio.NewMinioConnect,
		database.NewGormConnect,
	),
)

var ServiceModule = fx.Module("service",
	fx.Provide(
		JWTCONFIG,
		services.NewUserService,
		services.NewStorageService,
		services.NewJWTService,
		services.NewWebSocketService,
		services.NewMinioService,
	),
)

var ControllersModule = fx.Module("controllers",
	fx.Provide(controllers.NewFSController,
		controllers.NewUserController,
		controllers.NewWebsocketController,
	),
)

var RouteModule = fx.Module("route",
	fx.Provide(
		middlewares.NewAuth,
		router.InitRoute,
	),
)

type Params struct {
	fx.In
	Server *gin.Engine
	WS     *services.WebSocketService
}

func Hooks(lifecycle fx.Lifecycle, p Params) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				log.Println("启动文件同步服务端...")
				go p.WS.HandleWebSocketConnections()
				if err := p.Server.Run(":8080"); err != nil {
					panic(err)
				}
				return nil
			},
			OnStop: func(ctx context.Context) error {
				log.Println("关闭文件同步服务端...")
				return nil
			},
		},
	)
}

func JWTCONFIG() ([]byte, time.Duration, hash.Hash) {
	return []byte("zyl"), 24 * time.Hour, md5.New()
}
