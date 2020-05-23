package entity

import (
	"time"
)

// AccountForm ...
type AccountForm struct {
	PlatformID     string `json:"platform_id" form:"platform_id"`
	PlatformUserID string `json:"platform_user_id" form:"platform_user_id"`
	UserName       string `json:"user_name" form:"user_name"`
	Avatar         string `json:"avatar" form:"avatar"`
	Token          string `json:"token" form:"token"`
}

type SessionResp struct {
	Sign string `json:"sign"`
}

// SessionInfo ...
type SessionInfo struct {
	MinerID    uint64
	IP         string
	UA         string
	Phone      string
	InviteCode string
	LoginTime  time.Time
}

// USER
type ContextUser struct {
	PlatformID string `json:"pid"`
	UserID     string `json:"uid"`
	UserName   string `json:"name"`
	Avatar     string `json:"avatar"`
	UserType   string `json:"userType"`
}
