package service

import (
	"github.com/gomodule/redigo/redis"
	"sync"
	"time"
)

const (
	ProductStatusNormal = 0
	ProductStatusSaleOut = 1
	ProductStatusForceSaleOut = 2

)
type SecSkillConf struct {
	RedisBlackConf      RedisConf
	RedisProxy2LayerConf RedisConf
	EtcdConf       EtcdConf
	LogPath        string
	LogLevel       string
	SecProductInfoMap map[int]*SecProductInfoConf
	RWSecProductLock sync.RWMutex
	CookieSecretKey string
	UserSecAccessLimit int
	ReferWhiteList []string
	IPSecAccessLimit int
	ipBlackMap map[string]bool
	idBlackMap map[int]bool
	blackRedisPool *redis.Pool
	proxy2LayerRedisPool *redis.Pool
	secLimitMgr *SecLimitMgr
	RWBlackLock sync.RWMutex

	WriteProxy2LayerGoroutine_num int
	ReadProxy2LayerGoroutine_num int

	SecReqChan chan *SecRequest
	SecReqChanSize int
}

type SecProductInfoConf struct {
	ProductId  int
	StartTime  int64
	EndTime    int64
	Status     int
	TotalCount int
	LeftCount  int
}
type RedisConf struct {
	RedisAddr string
	RedisMaxIdle int
	RedisMaxActive  int
	RedisIdleTimeout int
}
type EtcdConf struct {
	EtcdAddr string
	Timeout int
	EtcdSecKeyPrefix string
	EtcdSecProductKey string
}

type SecRequest struct {
	ProductId int
	Source string
	AuthCode string
	SecTime string
	Nance string
	UserId int
	UserAuthSign string
	AccessTime time.Time
	ClientAddr string
	ClientRefence string


}
