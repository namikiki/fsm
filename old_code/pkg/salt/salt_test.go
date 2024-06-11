package salt

import (
	"crypto/md5"
	"log"
	"testing"
)

func TestNewSaltService(t *testing.T) {

	service := NewSaltService(md5.New())
	hashed := service.Hashed([]byte("123"), []byte("123"))
	log.Println(hashed)
}
