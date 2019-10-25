package middleware

import (
    "net/http"
    "github.com/gin-gonic/gin"
	"apigo/back/database"
    "apigo/back/models"
)

// IPFirewall : block banned IPs
func IPFirewall() gin.HandlerFunc { 
    return func(c *gin.Context) {
        r := c.Request
        ip := &models.Ip{}
        if !database.GetDB().Table("ips").Where("adress = ? AND blocked = ?", r.RemoteAddr, true).First(ip).RecordNotFound() { // check in the db if the ip is already registered and blocked
            c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
                "status":  http.StatusForbidden,
                "message": "Permission denied",
            })
            return
        }
    }
}