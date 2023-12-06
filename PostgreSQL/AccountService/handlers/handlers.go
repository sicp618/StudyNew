// handlers.go
package handlers

import (
	"net/http"

    "crypto/rand"
    "encoding/hex"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/gorilla/sessions"
	"gorm.io/gorm"
	"sicp618.com/hotpot/account/models"
)

type Handler struct {
	DB    *gorm.DB
	Pool  *redis.Client
	Store *sessions.CookieStore
}

func NewHandler(db *gorm.DB, pool *redis.Client, store *sessions.CookieStore) *Handler {
	return &Handler{DB: db, Pool: pool, Store: store}
}

func (h *Handler) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := user.IsValidUsername(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	user.Password = ""

	c.JSON(http.StatusOK, user)
}

func (h *Handler) Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var foundUser models.User
	if err := h.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&foundUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid email or password"})
		return
	}
	foundUser.Password = ""

    sessionToken := make([]byte, 32)
    rand.Read(sessionToken)

    // Store the session token and user ID in Redis
    err := h.Pool.Set(c, hex.EncodeToString(sessionToken), foundUser.ID, 0).Err()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"session-error": err.Error()})
        return
    }

    // Set the session token as a cookie
    http.SetCookie(c.Writer, &http.Cookie{
        Name:  "session_token",
        Value: hex.EncodeToString(sessionToken),
    })

	c.JSON(http.StatusOK, foundUser)
}

// return user info
func (h *Handler) User(c *gin.Context) {
    session, err := c.Request.Cookie("session_token")
    if err != nil {
        // c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        c.JSON(http.StatusUnauthorized, gin.H{"error": c.Request.Cookies()})
        return
    }

    // Get the user ID from Redis
    userID, err := h.Pool.Get(c, session.Value).Result()
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "No valid session"})
        return
    }

    // Get the user from the database
    var user models.User
    if err := h.DB.Where("id = ?", userID).First(&user).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    user.Password = ""

    c.JSON(http.StatusOK, user)
}
