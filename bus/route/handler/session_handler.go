package handler

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"oceanEngineService/bus/entity"
	"oceanEngineService/bus/entity/errmsg"
)

//ManagerLogin 登录
func GetSession(c *gin.Context) {
	var form entity.AccountForm
	if err := c.ShouldBindJSON(&form); err != nil {
		log.Error("parse parameter error")
		SendResponse(c, errmsg.ErrInvalidRequest, nil)
		return
	}
	//ip := getIP(c)
	//ua := getUA(c)
	if form.PlatformUserID == "" || form.UserName == "" || form.PlatformID == "" {
		log.Error("invalid parameter for request session")
		SendResponse(c, errmsg.ErrInvalidRequest, nil)
		return
	}
	//resp, err := service.LoginAccount(&form, 2)
	//SendResponseCheck(c, err, resp)
}
