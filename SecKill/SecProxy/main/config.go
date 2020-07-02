package main

import (
	"SecKill/SecProxy/service"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"strings"

)

var (
	secKillConf = &service.SecSkillConf{
		SecProductInfoMap : make(map[int]*service.SecProductInfoConf, 1024),
	}
)


func initConfig() (err error){
	redisAdrr := beego.AppConfig.String("redis_addr")
	//fmt.Printf("redisAdrr%v",redisAdrr)
	redisMaxIdle, _ := beego.AppConfig.Int("redis_max_idle")
	redisMaxActive, _ := beego.AppConfig.Int("redis_max_active")
	redisIdleTimeout, _ := beego.AppConfig.Int("redis_idle_timeout")
	etcdAddr := beego.AppConfig.String("etcd_addr")


	logs.Debug("read config succ, redis addr :%v", redisAdrr)
	logs.Debug("read config succ, etcd addr :%v", etcdAddr)

	secKillConf.EtcdConf.EtcdAddr = etcdAddr



	secKillConf.RedisConf.RedisAddr = redisAdrr
	secKillConf.RedisConf.RedisMaxIdle = redisMaxIdle
	secKillConf.RedisConf.RedisMaxActive = redisMaxActive
	secKillConf.RedisConf.RedisIdleTimeout = redisIdleTimeout
	if len(redisAdrr) == 0 || len(etcdAddr) == 0 {
		err = fmt.Errorf("init config failed, redis[%s] or etcd[%s] config is null", redisAdrr, etcdAddr)
	}

	etcdTimeout, err := beego.AppConfig.Int("etcd_timeout")
	fmt.Println("timeout", etcdTimeout)
	if err != nil {
		err = fmt.Errorf("init config failed, read etcd_timeout err :%v", err)
	}
	secKillConf.EtcdConf.Timeout = etcdTimeout
	secKillConf.EtcdConf.EtcdSecKeyPrefix = beego.AppConfig.String("etcd_sec_pre_key")
	//fmt.Println("prefix" , secKillConf.EtcdConf.EtcdSecKeyPrefix)
	if len(secKillConf.EtcdConf.EtcdSecKeyPrefix) == 0 {
		err = fmt.Errorf("init config failed, read etcd_sec_pre_ key err:%v ", err)
		return
	}
	productKey := beego.AppConfig.String("etcd_product_key")
	if len(productKey) == 0 {
		err = fmt.Errorf("init config failed, read etcd_product_key err: %v", err)
		return
	}
	if strings.HasSuffix(secKillConf.EtcdConf.EtcdSecKeyPrefix, "/") == false {
		secKillConf.EtcdConf.EtcdSecKeyPrefix = secKillConf.EtcdConf.EtcdSecKeyPrefix + "/"
	}
	secKillConf.EtcdConf.EtcdSecProductKey = fmt.Sprintf("%s%s", secKillConf.EtcdConf.EtcdSecKeyPrefix, productKey)
	//fmt.Println("etcdSecProductKey:",secKillConf.EtcdConf.EtcdSecProductKey)


	secKillConf.LogPath = beego.AppConfig.String("log_path")
	secKillConf.LogLevel = beego.AppConfig.String("log_level")
	//fmt.Println("logpath", secKillConf.LogPath)
	//fmt.Println("loglevel", secKillConf.LogLevel)

	secKillConf.CookieSecretKey = beego.AppConfig.String("cookie_secretkey")
	secLimit, err := beego.AppConfig.Int("user_sec_access_limit")
	if err != nil {
		err = fmt.Errorf("init config failed, read user_sec_access_limit err :%v", err)
		return
	}
	secKillConf.UserSecAccessLimit =secLimit
	referList := beego.AppConfig.String("refer_whitelist")
	if len(referList) > 0 {
		secKillConf.ReferWhiteList = strings.Split(referList, ",")
	}

	ipLimit, err := beego.AppConfig.Int("ip_sec_access_limit")
	if err != nil {
		err = fmt.Errorf("init config failed, read ip_sec_access_limit err :%v", err)
		return
	}
	secKillConf.IPSecAccessLimit = ipLimit
	return

}
