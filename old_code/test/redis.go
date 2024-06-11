package test

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"strconv"
)

var ctx = context.Background()

type MyObject struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestRedis() {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	objArray := []MyObject{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
	}

	for i, obj := range objArray {
		key := fmt.Sprintf("object:%d", i)
		err := rdb.HSet(ctx, key, map[string]interface{}{
			"name": obj.Name,
			"age":  obj.Age,
		}).Err()
		if err != nil {
			log.Fatalf("Error saving data to Redis: %v", err)
		}
	}

	for i := range objArray {
		key := fmt.Sprintf("object:%d", i)
		fields, err := rdb.HGetAll(ctx, key).Result()
		if err != nil {
			log.Fatalf("Error retrieving data from Redis: %v", err)
		}

		age, _ := strconv.Atoi(fields["age"])
		obj := MyObject{Name: fields["name"], Age: age}
		log.Printf("Retrieved object: %+v", obj)
	}

}
