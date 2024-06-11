package connect

import (
	"fsm/pkg/config"
	"log"
	"testing"

	"fsm/pkg/ent"
)

func TestT1(t *testing.T) {
	newConfig := config.NewConfig()
	connect := NewGormSQLiteConnect(newConfig)

	var res []ent.SyncTask

	connect.Debug().Model(&ent.SyncTask{}).Where(" user_id = ?", "u1").Find(&res)

	for s, s2 := range res {
		log.Println(s, s2)
	}
}
