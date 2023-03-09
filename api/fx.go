package api

import (
	"fsm/api/handle"

	"go.uber.org/fx"
)

var Module = fx.Module("api",
	handle.Module,
	fx.Provide(
		New,
	),
	fx.Invoke(AddRouters),
)
