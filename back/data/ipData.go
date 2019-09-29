package data

import (
	"apigo/back/database"
	"apigo/back/models"
	u "apigo/back/utils"

	"github.com/jinzhu/gorm"
)

//Everything that deals with ip, once the count is at least 3, the current ip is blocked

// Ip : ip model
type Ip models.Ip

// Increment ip
func (ip *Ip) Increment() map[string]interface{} {
	err := database.GetDB().Table("ips").Where(Ip{Adress: ip.Adress}).FirstOrCreate(ip).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Ip address not found")
		}
		return u.Message(false, "Connection error. Please retry")
	}
	ip.Count += 1
	if ip.Count >= 3 {
		ip.Blocked = true //ip is blocked
	}
	database.GetDB().Save(ip) //saving ip in database
	// Response
	return u.Message(false, "Invalid login credentials. Please try again")
}
