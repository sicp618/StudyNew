// handlers_test.go
package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/go-redis/redismock/v9"
	"github.com/redis/go-redis/v9"
	"sicp618.com/hotpot/account/models"
)

var db *gorm.DB
var store *sessions.CookieStore
var redisClient *redis.Client
var redisMock redismock.ClientMock

func init() {
    var err error
    db, err = gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    db.AutoMigrate(&models.User{})

	store = sessions.NewCookieStore([]byte("something-very-secret"))

	gin.SetMode(gin.TestMode)

	redisClient, redisMock = redismock.NewClientMock()
	redisMock.Regexp().ExpectSet(".*", ".*", 0).SetVal("OK")
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

	h := &Handler{DB: db, Store: store, Pool: redisClient}

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

	h := &Handler{DB: db, Store: store, Pool: redisClient}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request, _ = http.NewRequest(http.MethodGet, "/api/user/testuser2", nil)
	c.Params = []gin.Param{{Key: "username", Value: "testuser2"}}
	c.Request.Header.Set("Content-Type", "application/json")

	cookie := &http.Cookie{Name: "session_token", Value: fmt.Sprint(user.ID)}
	c.Request.AddCookie(cookie)
	redisMock.Regexp().ExpectGet(".*").SetVal(fmt.Sprintf("%d", user.ID))

	c.Set("current_user", *user)

	h.User(c)
	assert.Equal(t, http.StatusOK, w.Code)

	var returnedUser models.User
	err := json.Unmarshal(w.Body.Bytes(), &returnedUser)
	assert.Nil(t, err)
	assert.Equal(t, "testuser2", returnedUser.Username)
	assert.Equal(t, "", returnedUser.Password)
}

func TestAuthMiddleware(t *testing.T) {
	h := &Handler{DB: db, Store: store, Pool: redisClient}
    router := gin.Default()
    router.Use(h.AuthMiddleware())
    router.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, "Hello, World!")
    })

    req, err := http.NewRequest("GET", "/", nil)
	assert.Nil(t, err)

	// no session token
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)
    assert.Equal(t, http.StatusUnauthorized, w.Code)

	// invalid session token
	w = httptest.NewRecorder()
	redisMock.ClearExpect()
	redisMock.ExpectGet("124").SetVal("100")
	cookie := &http.Cookie{Name: "session_token", Value: "123"}
	req.AddCookie(cookie)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	// user not found
	w = httptest.NewRecorder()
	redisMock.ClearExpect()
	redisMock.ExpectGet("123").SetVal("100")
	req.AddCookie(cookie)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)

	// normal
	user := &models.User{Username: "testuser5", Password: "test123456", Email: "testuser5@email.com"}
	h.DB.Create(user)
	w = httptest.NewRecorder()
	redisMock.ClearExpect()
	redisMock.ExpectGet("123").SetVal(fmt.Sprint(user.ID))
	req.AddCookie(cookie)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}