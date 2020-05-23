package handler

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"oceanEngineService/bus/entity"
	"oceanEngineService/bus/entity/errmsg"
)

// SendStatusResponse ...
func SendStatusResponse(c *gin.Context, status int, err errmsg.Error, data interface{}) {
	log.WithFields(log.Fields{
		"code": status,
		"msg":  err,
		"data": data,
	}).Info("hi")

	c.JSON(status, entity.Response{
		Code: err.Code(),
		Msg:  err.Error(),
		Data: data,
	})
	c.Abort()
}

// SendResponse ...
func SendResponse(c *gin.Context, err errmsg.Error, data interface{}) {
	SendStatusResponse(c, http.StatusOK, err, data)
}

// SendResponseCheck ...
func SendResponseCheck(c *gin.Context, err error, data interface{}) {
	if nil != err {
		SendResponse(c, errmsg.NewInternalErr(err), nil)
		return
	}
	SendResponse(c, errmsg.NoErr, data)
}
