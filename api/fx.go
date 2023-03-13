package api

import (
	"fsm/api/handle"

	"github.com/go-playground/validator/v10"
	"go.uber.org/fx"
)

var Module = fx.Module("api",
	//handle.Module,
	fx.Provide(
		validator.New,
		handle.NewUser,
		handle.NewFile,
		handle.NewDir,
		handle.NewCommon,
		handle.NewSyncTask,
		New,
	),
	fx.Invoke(AddRoutes),
)
