// models_test.go
package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidUsername(t *testing.T) {
	tests := []struct {
		username string
		wantErr  bool
	}{
		{"validusername", false},
		{"in valid", true},
		{"toolongusername............................................................", true},
		{"short", false},
		{"with/slash", true},
		{"with?question", true},
		{"with#hash", true},
	}

	for _, tt := range tests {
		t.Run(tt.username, func(t *testing.T) {
			user := &User{Username: tt.username}
			err := user.IsValidUsername()
			if tt.wantErr {
				assert.Error(t, err, tt.username)
			} else {
				assert.NoError(t, err, tt.username)
			}
		})
	}
}