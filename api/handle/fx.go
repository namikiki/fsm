package handle

import (
	"github.com/go-playground/validator/v10"
	"go.uber.org/fx"
)

var Module = fx.Module("handler",
	fx.Provide(
		validator.New,
		NewUser,
		NewFile,
		NewDir,
		NewCommon,
		NewSyncTask,
	),
)
