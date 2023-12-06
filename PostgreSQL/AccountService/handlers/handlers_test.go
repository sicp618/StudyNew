// handlers_test.go
package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"sicp618.com/hotpot/account/models"
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
	h := &Handler{DB: db}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request, _ = http.NewRequest(http.MethodPost, "/register11", strings.NewReader(`{"username":"testuser","password":"testpass","email":"testuser@email.com"}`))
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

	h := &Handler{DB: db, Store: store}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req, _ := http.NewRequest(http.MethodPost, "/api/login", strings.NewReader(`{"email":"test1@test.com","password":"testpass"}`))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	assert.NotNil(t, c.Request)
	assert.NotNil(t, h)
	h.Login(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var returnedUser models.User
	err := json.Unmarshal(w.Body.Bytes(), &returnedUser)
	assert.Nil(t, err)
	assert.Equal(t, "testuser1", returnedUser.Username)
	assert.Equal(t, "", returnedUser.Password)
}

func TestUser(t *testing.T) {
	user := &models.User{Username: "testuser2", Password: "test123456", Email: "testuser2@email.com"}
	db.Create(user)

	h := &Handler{DB: db}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request, _ = http.NewRequest(http.MethodGet, "/api/user/testuser2", nil)
	c.Params = []gin.Param{{Key: "username", Value: "testuser2"}}
	c.Request.Header.Set("Content-Type", "application/json")

	h.User(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var returnedUser models.User
	err := json.Unmarshal(w.Body.Bytes(), &returnedUser)
	assert.Nil(t, err)
	assert.Equal(t, "testuser2", returnedUser.Username)
	assert.Equal(t, "", returnedUser.Password)
}
