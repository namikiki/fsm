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

// RepositoriesModule 定义了存储库模块，提供 StorageRepository 和 UserRepository
var RepositoriesModule = fx.Module("RepositoriesModule",
	fx.Provide(
		repositories.NewStorageRepository,
		repositories.NewUserRepository,
	),
)

// PKGModule 定义了基本包模块，提供 Config、Redis、Minio 连接和 Gorm 连接
var PKGModule = fx.Module("pkg",
	fx.Provide(
		configs.NewConfig,
		redis.NewRedis,
		minio.NewMinioConnect,
		database.NewGormConnect,
	),
)

// ServiceModule 定义了服务模块，提供各种服务
var ServiceModule = fx.Module("service",
	fx.Provide(
		JWTCONFIG,
		services.NewUserService,
		services.NewFolderService,
		services.NewFileService,
		services.NewJWTService,
		services.NewWebSocketService,
		services.NewMinioService,
	),
)

// ControllersModule 定义了控制器模块，提供各种控制器
var ControllersModule = fx.Module("controllers",
	fx.Provide(controllers.NewFolderController,
		controllers.NewFileController,
		controllers.NewUserController,
		controllers.NewWebsocketController,
	),
)

// RouteModule 定义了路由模块，提供认证中间件和路由初始化
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

// Hooks 定义应用的生命周期钩子
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

// JWTCONFIG 配置 JWT，返回密钥、过期时间和哈希函数
func JWTCONFIG() ([]byte, time.Duration, hash.Hash) {
	return []byte("zyl"), 24 * time.Hour, md5.New()
}
