package middleware

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"oceanEngineService/bus/entity/errmsg"
	"oceanEngineService/bus/route/handler"
	"oceanEngineService/core/pkg"
)

func AuthMiddleware(accountType int) gin.HandlerFunc {
	return func(c *gin.Context) {
		cxt, err := pkg.ParseRequest(c, accountType)
		if err != nil {
			log.Errorf("AuthMiddleware error: %s\n", err)
			handler.SendResponse(c, errmsg.ErrTokenInvalidRequest, nil)
			c.Abort()
			return
		}
		c.Set(pkg.UserName, cxt.Username)
		c.Set(pkg.Avatar, cxt.Avatar)
		c.Set(pkg.UID, cxt.UID)
		c.Set(pkg.PID, cxt.PID)
		c.Set(pkg.UserType, cxt.UserType)

		c.Next()
	}
}
