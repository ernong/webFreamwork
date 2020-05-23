package pkg

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

var (
	ErrMissingHeader = errors.New("the length of the authorization header is zero")
)

type Context struct {
	PID      string //平台id
	UID      string //平台用户id
	Username string //平台用户名
	Avatar   string //平台用户头像
	UserType string //用户类型
}

func secretFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(secret), nil
	}
}

func Parse(tokenString string, secret string) (*Context, error) {
	ctx := &Context{}

	token, err := jwt.Parse(tokenString, secretFunc(secret))

	if err != nil {
		return ctx, err

	} else if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		ctx.UID = claims["uid"].(string)
		ctx.PID = claims["pid"].(string)
		ctx.Username = claims["username"].(string)
		ctx.Avatar = claims["avatar"].(string)
		ctx.UserType = claims["userType"].(string)
		return ctx, nil

	} else {
		return ctx, err
	}
}

func ParseRequest(c *gin.Context, accountType int) (*Context, error) {
	header := c.Request.Header.Get("Authorization")

	secret := ""
	if accountType == 1 {
		secret = viper.GetString("jwt_secret_admin")
	} else {
		secret = viper.GetString("jwt_secret_user")
	}

	if len(header) == 0 {
		return &Context{}, ErrMissingHeader
	}

	var t string

	fmt.Sscanf(header, "Backend_%s", &t)
	return Parse(t, secret)
}

func Sign(c Context, secret string, accountType int) (tokenString string, err error) {
	if accountType == 1 {
		secret = viper.GetString("jwt_secret_admin")
	} else {
		secret = viper.GetString("jwt_secret_user")
	}
	ts := time.Now().Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"pid":      c.PID,
		"uid":      c.UID,
		"username": c.Username,
		"avatar":   c.Avatar,
		"userType": fmt.Sprintf("%v", accountType),
		"exp":      ts + 60*30,
		"nbf":      ts,
		"iat":      ts,
	})

	tokenString, err = token.SignedString([]byte(secret))
	log.Debugf("token:%v, ctime:%v,expTime:%v", tokenString, ts, ts+60*30)
	return
}
