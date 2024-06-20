package database

import (
	"fmt"
	"fsm/models"
	"fsm/pkg/types"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewGormConnect GORM数据库连接
// todo 多数据库支持
func NewGormConnect(config *types.Config) *gorm.DB {
	// 拼接成 MySQL DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DataBase.User,
		config.DataBase.Password,
		config.DataBase.Host,
		config.DataBase.Port,
		config.DataBase.BDName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库 连接错误:" + err.Error())
	}

	//db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	//if err != nil {
	//	panic(err)
	//}

	//db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	//if err != nil {
	//	panic(err)
	//}
	if err := db.AutoMigrate(&models.User{}, &models.Folder{}, &models.File{}); err != nil {
		panic(err)
	}

	return db
}
