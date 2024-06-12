package main

import (
	"context"
	"fsm/di"
	"go.uber.org/fx"
	"log"
)

func main() {

	// 创建一个新的 fx 应用
	app := fx.New(
		di.PKGModule,          // 引入基本包模块
		di.ServiceModule,      // 引入服务模块
		di.ControllersModule,  // 引入控制器模块
		di.RepositoriesModule, // 引入存储库模块
		di.RouteModule,        // 引入路由模块
		fx.Invoke(di.Hooks),   // 调用 Hooks 函数，定义应用的生命周期钩子
	)

	// 启动 fx 应用
	if err := app.Start(context.Background()); err != nil {
		log.Fatal(err) // 如果启动失败，记录错误日志并终止程序
	}

	// 停止 fx 应用
	if err := app.Stop(context.Background()); err != nil {
		log.Fatal(err) // 如果停止失败，记录错误日志并终止程序
	}
}
