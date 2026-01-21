package constants

import "time"

const (
	MaxConnections  = 500              // (DB) 最大连接数
	MaxIdleConns    = 200              // (DB) 最大空闲连接数
	ConnMaxLifetime = 1 * time.Hour    // (DB) 最大可复用时间
	ConnMaxIdleTime = 30 * time.Minute // (DB) 最长保持空闲状态时间
)
