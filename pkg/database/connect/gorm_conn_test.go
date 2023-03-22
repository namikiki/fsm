package connect

import (
	"log"
	"testing"

	"fsm/pkg/ent"
)

func TestT1(t *testing.T) {
	connect := NewGormSQLiteConnect()

	//res := make(map[string]string)

	var res []ent.SyncTask

	connect.Debug().Model(&ent.SyncTask{}).Where(" user_id = ?", "u1").Find(&res)

	for s, s2 := range res {
		log.Println(s, s2)
	}
}
