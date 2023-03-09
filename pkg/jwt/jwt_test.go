package jwt

import (
	"context"
	"log"
	"testing"
	"time"
)

func TestNewJWTService(t *testing.T) {

	key := []byte("1233")
	uuid := "zyl"

	claims := NewJWTService(key, 24*time.Hour)
	gen, err := claims.Gen(context.Background(), uuid)
	if err != nil {
		t.Fatalf("生成jwt失败: %v", err)
	}
	log.Println(gen)

	pa, err := claims.Parse(context.Background(), gen)
	if err != nil {
		t.Fatalf("解析jwt失败: %v", err)
	}
	log.Println(pa)
}
