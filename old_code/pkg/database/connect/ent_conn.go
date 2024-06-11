package connect

import (
	"context"
	"log"

	"fsm/pkg/ent"

	_ "github.com/mattn/go-sqlite3"
)

func NewEntSQLiteConnect() *ent.Client {
	conn, err := ent.Open("sqlite3", "file:ent.db?mode=rwc&cache=shared&_fk=1&_cache_size=20000")
	if err != nil {
		log.Println(err)
	}

	// 自动迁移 ent/schema 对象结构
	if err := conn.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return conn
}
