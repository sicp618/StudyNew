// handlers_test.go
package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
    "github.com/gorilla/sessions"

	"sicp618.com/hotpot/account/models"
	"sicp618.com/hotpot/account/handlers"
)

var db *gorm.DB
var store *sessions.CookieStore

func init() {
    var err error
    db, err = gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    db.AutoMigrate(&models.User{})

	store = sessions.NewCookieStore([]byte("something-very-secret"))

	gin.SetMode(gin.TestMode)
}

func TestRegister(t *testing.T) {
	h := &handlers.Handler{DB: db}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request, _ = http.NewRequest(http.MethodPost, "/register", strings.NewReader(`{"username":"testuser","password":"testpass","email":"testuser@email.com"}`))
	c.Request.Header.Set("Content-Type", "application/json")

	h.Register(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var user models.User
	db.Where("email = ?", "testuser@email.com").First(&user)
	assert.Equal(t, "testuser", user.Username)
}

func TestLogin(t *testing.T) {
	user := &models.User{Username: "testuser1", Password: "testpass", Email: "test1@test.com"}
	db.Create(user)

	h := &handlers.Handler{DB: db, Store: store}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request, _ = http.NewRequest(http.MethodPost, "/login", strings.NewReader(`{"email":"test1@test.com","password":"testpass"}`))
	c.Request.Header.Set("Content-Type", "application/json")
	assert.NotNil(t, c.Request)
	assert.NotNil(t, h)
	h.Login(c)

	assert.Equal(t, http.StatusOK, w.Code)
}