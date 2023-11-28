// models.go
package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"username"`
	Password string `json:"-"`
	Email    string `gorm:"unique;not null" json:"email"`
	Avatar   string `json:"avatar"`
	Bio      string `json:"bio"`
}
