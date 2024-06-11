package connect

import (
	"fmt"
	"fsm/pkg/ent"
	"fsm/pkg/types"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewGormSQLiteConnect GORM数据库连接
// todo 多数据库支持
func NewGormSQLiteConnect(config *types.Config) *gorm.DB {
	// 拼接成 MySQL DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DataBase.User,
		config.DataBase.Password,
		config.DataBase.Host,
		config.DataBase.Port,
		config.DataBase.BDName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	//db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	//if err != nil {
	//	panic(err)
	//}

	//db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	//if err != nil {
	//	panic(err)
	//}

	if err := db.AutoMigrate(&ent.User{}, &ent.SyncTask{}, &ent.Dir{}, &ent.File{}); err != nil {
		panic(err)
	}

	return db
}
