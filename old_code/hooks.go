package old_code

import (
	"context"
	"fsm/pkg/sync"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"log"
)

type Params struct {
	fx.In
	Server *gin.Engine
	Syncer *sync.Syncer
}

func registerHooks(lifecycle fx.Lifecycle, p Params) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				log.Println("启动文件同步服务端...")
				go p.Syncer.WebSocketLoop()
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
