package middleware

import (
    "net/http"
    "github.com/gin-gonic/gin"
	"apigo/back/database"
	"apigo/back/models"
)

// IPFirewall : block banned IPs
func IPFirewall() gin.HandlerFunc {
    ip := &models.Ip{}
    return func(c *gin.Context) {
        if !database.GetDB().Table("users").Where("address = ? AND blocked = ?", c.ClientIP(), true).First(ip).RecordNotFound() {
            c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
                "status":  http.StatusForbidden,
                "message": "Permission denied",
            })
            return
        }
    }
}