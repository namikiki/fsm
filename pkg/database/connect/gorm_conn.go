package connect

import (
	"fsm/pkg/ent"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewGormSQLiteConnect() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&ent.User{}, &ent.SyncTask{}, &ent.Dir{}, &ent.File{}); err != nil {
		panic(err)
	}

	return db
}
