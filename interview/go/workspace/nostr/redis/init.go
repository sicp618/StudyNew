package redis

import (
    "context"
    "fmt"
    "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

var rdb *redis.Client

func Init() error {
    fmt.Println("Connecting to Redis...")

    rdb = redis.NewClient(&redis.Options{
        Addr:     "localhost:6379", // Redis 服务器地址
        Password: "",               // 没有密码时为空字符串
        DB:       0,                // 使用默认数据库
    })

    // 测试连接
    pong, err := rdb.Ping(ctx).Result()
    if err != nil {
        fmt.Println("Failed to connect to Redis:", err)
        return err
    }
    fmt.Println("Connected to Redis:", pong)

    return nil
}

func Close() {
    fmt.Println("Closing Redis connection...")
    if err := rdb.Close(); err != nil {
        fmt.Println("Failed to close Redis connection:", err)
    }
}

func Read() {
    // 读取数据
    val, err := rdb.Get(ctx, "key").Result()
    if err == redis.Nil {
        fmt.Println("Key does not exist")
    } else if err != nil {
        fmt.Println("Error getting key:", err)
    } else {
        fmt.Println("Key value:", val)
    }
}