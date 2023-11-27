// models.go
package models

import "github.com/jinzhu/gorm"

type User struct {
    gorm.Model
    Username    string `gorm:"unique;not null"`
    Email       string `gorm:"unique;not null"`
    Password    string `gorm:"not null"`
    Intro       string
    Avatar      string
}