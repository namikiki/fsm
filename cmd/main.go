package main

import (
	"context"
	"fsm/di"
	"go.uber.org/fx"
	"log"
)

func main() {

	app := fx.New(
		di.PKGModule,
		di.ServiceModule,
		di.ControllersModule,
		di.RepositoriesModule,
		di.RouteModule,
		fx.Invoke(di.Hooks),
	)

	if err := app.Start(context.Background()); err != nil {
		log.Fatal(err)
	}

	if err := app.Stop(context.Background()); err != nil {
		log.Fatal(err)
	}
}
