package service

import "sync"

const (
	ProductStatusNormal = 0
	ProductStatusSaleOut = 1
	ProductStatusForceSaleOut = 2
)
type SecSkillConf struct {
	RedisConf      RedisConf
	EtcdConf       EtcdConf
	LogPath        string
	LogLevel       string
	SecProductInfoMap map[int]*SecProductInfoConf
	RWSecProductLock sync.RWMutex
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

