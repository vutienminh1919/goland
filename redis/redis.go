package redis

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var Rdb *redis.Client
var Ctx = context.Background()

func ConnectRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // tùy chỉnh port nếu cần
		Password: "",               // nếu có password thì thêm
		DB:       0,                // DB mặc định
	})

	// Test kết nối
	_, err := Rdb.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("❌ Kết nối Redis thất bại: %v", err)
	}

	log.Println("✅ Đã kết nối Redis thành công")
}
