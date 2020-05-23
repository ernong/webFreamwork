package handler

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net"
	"oceanEngineService/bus/entity"
	"oceanEngineService/bus/entity/errmsg"
	"time"
)

// variable for session
var (
	SessionLifeTime = time.Duration(24 * time.Hour)
)

func getIP(c *gin.Context) string {
	ip := c.ClientIP()

	realIP := c.GetHeader("X-REAL-IP")
	ipAddr := net.ParseIP(realIP)
	if nil != ipAddr {
		return ipAddr.String()
	}
	return ip
}

func getUA(c *gin.Context) string {
	return c.GetHeader("User-Agent")
}

func getLogger(c *gin.Context) *log.Entry {
	return log.WithFields(log.Fields{
		"IP":        c.ClientIP(),
		"UA":        c.GetHeader("User-Agent"),
		"X-Real-IP": getIP(c),
	})
}

func clearSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
}

func getSession(c *gin.Context) (*entity.SessionInfo, error) {
	session := sessions.Default(c)
	infoRaw := session.Get("info")
	if nil == infoRaw {
		return nil, errmsg.ErrRequireLogin
	}
	info, ok := session.Get("info").(entity.SessionInfo)
	if !ok {
		return nil, errmsg.ErrInternalServer
	}

	if time.Now().Sub(info.LoginTime) > SessionLifeTime {
		log.Error("session expired")
		return nil, errmsg.ErrRequireLogin
	}
	return &info, nil
}

func verifySession(c *gin.Context) (*entity.SessionInfo, error) {
	info, err := getSession(c)
	if nil != err {
		SendResponse(c, errmsg.ErrRequireLogin, nil)
		return nil, err
	}
	return info, nil
}
