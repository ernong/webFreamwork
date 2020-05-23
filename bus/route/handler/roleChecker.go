package handler

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"oceanEngineService/bus/entity"
)

func IsAdmin(c *gin.Context) bool {
	_ = GetContextInfo(c)
	//if user.UserType == fmt.Sprintf("%v", db.UserTypeAdmin) ||
	//	user.UserType == fmt.Sprintf("%v", db.UserTypeManager) {
	//	return true
	//}
	return false
}

func IsUser(c *gin.Context) bool {
	_ = GetContextInfo(c)
	//if user.UserType == fmt.Sprintf("%v", db.UserTypeUser) {
	//	return true
	//}
	return false
}

func IsManager(c *gin.Context) bool {
	_ = GetContextInfo(c)
	//if user.UserType == fmt.Sprintf("%v", db.UserTypeManager) {
	//	return true
	//}
	return false
}

func IsSuperAdmin(c *gin.Context) bool {
	_ = GetContextInfo(c)
	//if user.UserType == fmt.Sprintf("%v", db.UserTypeAdmin) {
	//	return true
	//}
	return false
}

func GetContextInfo(c *gin.Context) *entity.ContextUser {
	if c == nil {
		return nil
	}
	user := &entity.ContextUser{}
	if name, ok := c.Get("username"); ok {
		if name != nil {
			user.UserName = name.(string)
		}
	} else {
		log.Debugf("get username in context err")
	}

	if userType, ok := c.Get("userType"); ok {
		if userType != nil {
			user.UserType = userType.(string)
		}
	} else {
		log.Debugf("get userType in context err")
	}
	return user
}
