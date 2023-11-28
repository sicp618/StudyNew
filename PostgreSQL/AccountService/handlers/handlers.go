// handlers.go
package handlers

import (
    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    "github.com/gorilla/sessions"
    "github.com/go-redis/redis"
    "net/http"
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

    if err := h.DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": user})
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

    session, _ := h.Store.Get(c.Request, "session-name")
    session.Values["user_id"] = foundUser.ID
    session.Save(c.Request, c.Writer)

    c.JSON(http.StatusOK, gin.H{"data": foundUser})
}