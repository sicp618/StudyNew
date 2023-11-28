// models.go
package models

import (
	"errors"
	"net/url"
	"unicode"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"username"`
	Password string `json:"-"`
	Email    string `gorm:"unique;not null" json:"email"`
	Avatar   string `json:"avatar"`
	Bio      string `json:"bio"`
}

func (u *User) IsValidUsername() error {
	if len(u.Username) < 3 {
		return errors.New("username must be at least 3 characters long")
	}

	if len(u.Username) > 64 {
		return errors.New("username must be no more than 64 characters long")
	}

	for _, char := range u.Username {
		if !unicode.IsLetter(char) && !unicode.IsNumber(char) && char != '-' && char != '_' && char != '.' && char != '~' {
			return errors.New("username contains invalid characters")
		}
	}

	if u.Username != url.PathEscape(u.Username) {
		return errors.New("username must be URL safe")
	}

	return nil
}