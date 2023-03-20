package main

import (
	"crypto/md5"
	"hash"
	"log"
	"os"
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

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/fx"
)

func Init() ([]byte, time.Duration, hash.Hash) {
	return []byte("zyl"), 24 * time.Hour, md5.New()
}

func main() {
	err := os.Remove("gorm.db")
	if err != nil {
		log.Println(err)
	}

	var server *gin.Engine
	var syncer *sync.Syncer

	if err := fx.New(

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
		//fx.Populate(&server, &websocketPool),
		fx.Populate(&server, &syncer),
	).Err(); err != nil {
		panic(err)
	}

	go syncer.WebSocketLoop()
	if err := server.Run(":8080"); err != nil {
		panic(err)
	}
}
