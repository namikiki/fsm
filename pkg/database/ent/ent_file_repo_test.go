package ent

import "testing"

//import (
//	"context"
//	"log"
//	"testing"
//	"time"
//
//	"fsm/pkg/database/connect"
//	"fsm/pkg/domain"
//	"fsm/pkg/ent"
//
//	"github.com/stretchr/testify/assert"
//)
//
//func initCase() (domain.FileRepository, ent.File) {
//	repository := NewMysqlFileRepository(connect.NewEntSQLiteConnect())
//	file := ent.File{
//		ID:         "di1231",
//		ParentName: "123123",
//		UserID:     "1231231",
//		Name:       "1231231",
//		Size:       12,
//		Deleted:    false,
//		CreateTime: time.Now(),
//		ModTime:    time.Now(),
//	}
//	return repository, file
//}
//
//func TestFileStore(t *testing.T) {
//	repository, file := initCase()
//	err := repository.Store(context.Background(), file)
//	if err != nil {
//		log.Printf("err %v", err)
//	}
//}
//
//func TestFileRemove(t *testing.T) {
//	repository, file := initCase()
//	err := repository.Remove(context.Background(), file.ID, file.UserID)
//	if err != nil {
//		log.Printf("err %v", err)
//	}
//}

func TestFileGet(t *testing.T) {
	//repository, file := initCase()
	//f, err := repository.Get(context.Background(), file.ID, file.UserID)
	//if err != nil {
	//	log.Printf("err %v", err)
	//}
	//assert.Equal(t, file.ID, f.ID, "they should be equal")
	//assert.Equal(t, file.UserID, f.UserID, "they should be equal")
}

func TestFileUpdate(t *testing.T) {
	//repository, file := initCase()
	//err := repository.Update(context.Background(), file.ID, file.UserID, 299, time.Now())
	//if err != nil {
	//	log.Printf("err %v", err)
	//}

}
