package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
	"gorm.io/gorm"    
	"gorm.io/driver/postgres"    
    "github.com/go-redis/redis"
    "github.com/gorilla/sessions"
    "github.com/spf13/viper"
    "sicp618.com/hotpot/account/handlers"
    "sicp618.com/hotpot/account/models"
    _ "sicp618.com/hotpot/account/store"
)

func readConfig() {
    viper.SetConfigName("config")
    viper.AddConfigPath(".")
    err := viper.ReadInConfig()
    if err != nil {
        panic(fmt.Errorf("fatal error config file: %s", err))
    }
}

func initDB() *gorm.DB {
    dsn := viper.GetString("dsn")
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic(fmt.Sprintln("failed to connect database", err))
    }

	db.AutoMigrate(&models.User{})

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
    r.GET("/api/user/:username", handlers.User)

    return r
}

func main() {
	readConfig()
    db := initDB()

    client := initRedis()
	defer client.Close()

    r := initService(db, client)
    r.Run() 
}