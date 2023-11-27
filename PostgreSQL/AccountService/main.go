// main.go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/sessions"
	_ "github.com/jinzhu/gorm/dialects/postgres"
    "github.com/spf13/viper"
	"sicp618.com/hotpot/account/handlers"
	"sicp618.com/hotpot/account/models"
	"sicp618.com/hotpot/account/store"
)

func main() {
    viper.SetConfigName("config")
    viper.AddConfigPath(".")
    err := viper.ReadInConfig()
    if err != nil {
        panic(fmt.Errorf("Fatal error config file: %s", err))
    }

    dsn := viper.GetString("dsn")

    redisAddress := viper.GetString("redis.address")
    redisPassword := viper.GetString("redis.password")
    redisDB := viper.GetInt("redis.db")

	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	redisPool := &redis.Pool{
		Dial: func () (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}

	store := sessions.NewCookieStore([]byte("something-very-secret"))

	r := gin.Default()

	handlers := handlers.NewHandler(db, redisPool, store)
	
	r.GET("/api/resource", handlers.GetResource)
	r.POST("/api/resource", handlers.CreateResource)
	r.PUT("/api/resource/:id", handlers.UpdateResource)
	r.DELETE("/api/resource/:id", handlers.DeleteResource)

	r.Run() // listen and serve on 0.0.0.0:8080
}