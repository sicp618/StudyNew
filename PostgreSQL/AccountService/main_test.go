package main

import (
	"testing"

    "github.com/spf13/viper"
	"gorm.io/driver/sqlite"
)

func init() {
    dsn := sqlite.Open("file::memory:")
	viper.Set("dsn", dsn)
}

func TestInitRedis(t *testing.T) {
}

func TestInitService(t *testing.T) {
}