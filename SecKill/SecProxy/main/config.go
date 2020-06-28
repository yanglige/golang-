package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

var (
	secKillConf = &SecSkillConf{}
)


type SecSkillConf struct {
	redisConf RedisConf
	etcdConf EtcdConf
	logPath string
	logLevel string
}

type RedisConf struct {
	redisAddr string
	redisMaxIdle int
	redisMaxActive  int
	redisIdleTimeout int
}
type EtcdConf struct {
	etcdAddr string
	timeout int
}

func initConfig() (err error){
	redisAdrr := beego.AppConfig.String("redis_addr")
	redisMaxIdle, _ := beego.AppConfig.Int("redis_max_idle")
	redisMaxActive, _ := beego.AppConfig.Int("redis_max_active")
	redisIdleTimeout, _ := beego.AppConfig.Int("redis_idle_timeout")
	etcdAddr := beego.AppConfig.String("etcd_addr")


	logs.Debug("read config succ, redis addr :%v", redisAdrr)
	logs.Debug("read config succ, etcd addr :%v", etcdAddr)

	secKillConf.etcdConf.etcdAddr = etcdAddr
	secKillConf.redisConf.redisAddr = redisAdrr
	secKillConf.redisConf.redisMaxIdle = redisMaxIdle
	secKillConf.redisConf.redisMaxActive = redisMaxActive
	secKillConf.redisConf.redisIdleTimeout = redisIdleTimeout
	if len(redisAdrr) == 0 || len(etcdAddr) == 0 {
		err = fmt.Errorf("init config failed, redis[%s] or etcd[%s] config is null", redisAdrr, etcdAddr)
	}

	etcdTimeout, err := beego.AppConfig.Int("etcd_timeout")
	if err != nil {
		err = fmt.Errorf("init config failed, read etcd_timeout err = %v", err)
	}
	secKillConf.etcdConf.timeout = etcdTimeout

	secKillConf.logPath = beego.AppConfig.String("log_path")
	secKillConf.logLevel = beego.AppConfig.String("log_level")



	return
}
