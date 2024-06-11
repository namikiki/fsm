package main

import (
	"context"
	"crypto/md5"
	"hash"
	"log"
	"time"

	"fsm/api"
	"fsm/pkg/config"
	"fsm/pkg/database/connect"
	"fsm/pkg/database/gorm"
	"fsm/pkg/jwt"
	"fsm/pkg/redis"
	"fsm/pkg/salt"
	"fsm/pkg/storage/minio"
	"fsm/pkg/sync"
	"fsm/pkg/user"

	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/fx"
)

func Init() ([]byte, time.Duration, hash.Hash) {
	return []byte("zyl"), 24 * time.Hour, md5.New()
}

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	app := fx.New(

		fx.Provide(
			Init,
			config.NewConfig,
			redis.NewRedis,
			minio.NewMinioFileStorageService,
			minio.NewMinioConnect,
		),

		fx.Provide( // 数据库
			//connect.NewEntSQLiteConnect,
			//ent.NewMysqlUserRepository,
			//ent.NewMysqlFileRepository,

			connect.NewGormSQLiteConnect,
			gorm.NewDirRepository,
			gorm.NewUserRepository,
			gorm.NewFileRepository,
			gorm.NewSyncRepository,
		),

		fx.Provide(
			sync.NewSyncer,
			user.NewService,
			jwt.NewJWTService,
			salt.NewSaltService,
		),

		api.Module,
		fx.Invoke(registerHooks),
	)

	if err := app.Start(context.Background()); err != nil {
		log.Fatal(err)
	}

	if err := app.Stop(context.Background()); err != nil {
		log.Fatal(err)
	}
}
