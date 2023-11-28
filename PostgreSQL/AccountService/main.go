package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    "github.com/go-redis/redis"
    "github.com/gorilla/sessions"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "github.com/spf13/viper"
    "sicp618.com/hotpot/account/handlers"
    _ "sicp618.com/hotpot/account/models"
    _ "sicp618.com/hotpot/account/store"
)

func initDB() *gorm.DB {
    viper.SetConfigName("config")
    viper.AddConfigPath(".")
    err := viper.ReadInConfig()
    if err != nil {
        panic(fmt.Errorf("Fatal error config file: %s", err))
    }

    dsn := viper.GetString("dsn")

    db, err := gorm.Open("postgres", dsn)
    if err != nil {
        panic(fmt.Sprintln("failed to connect database", err))
    }
    return db
}

func initRedis() *redis.Client {
    redisAddress := viper.GetString("redis.address")
    redisPassword := viper.GetString("redis.password")
    redisDB := viper.GetInt("redis.db")

    client := redis.NewClient(&redis.Options{
        Addr:     redisAddress,
        Password: redisPassword,
        DB:       redisDB,
    })

    pong, err := client.Ping().Result()
    if err != nil {
        panic(err)
    } else {
        fmt.Println(pong)
    }
    return client
}

func initService(db *gorm.DB, client *redis.Client) *gin.Engine {
    store := sessions.NewCookieStore([]byte("something-very-secret"))

    r := gin.Default()

    handlers := handlers.NewHandler(db, client, store)
    
    r.POST("/api/login", handlers.Login)
    r.POST("/api/register", handlers.Register)

    return r
}

func main() {
    db := initDB()
    defer db.Close()

    client := initRedis()

    r := initService(db, client)

    r.Run() 
}