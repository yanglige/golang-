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
	redisBlackAdrr := beego.AppConfig.String("redis_black_addr")
	//fmt.Printf("redisAdrr%v",redisAdrr)
	redisMaxIdle, _ := beego.AppConfig.Int("redis_max_idle")
	redisMaxActive, _ := beego.AppConfig.Int("redis_max_active")
	redisIdleTimeout, _ := beego.AppConfig.Int("redis_idle_timeout")
	etcdAddr := beego.AppConfig.String("etcd_addr")


	logs.Debug("read config succ, redis addr :%v", redisBlackAdrr)
	logs.Debug("read config succ, etcd addr :%v", etcdAddr)

	secKillConf.EtcdConf.EtcdAddr = etcdAddr



	secKillConf.RedisBlackConf.RedisAddr = redisBlackAdrr
	secKillConf.RedisBlackConf.RedisMaxIdle = redisMaxIdle
	secKillConf.RedisBlackConf.RedisMaxActive = redisMaxActive
	secKillConf.RedisBlackConf.RedisIdleTimeout = redisIdleTimeout
	if len(redisBlackAdrr) == 0 || len(etcdAddr) == 0 {
		err = fmt.Errorf("init config failed, redis[%s] or etcd[%s] config is null", redisBlackAdrr, etcdAddr)
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





	redisProxy2LayerAdrr := beego.AppConfig.String("redis_proxy2layer_addr")
	//fmt.Printf("redisAdrr%v",redisAdrr)
	redisProxy2LayerIdle, _ := beego.AppConfig.Int("redis_proxy2layer_idle")
	redisProxy2LayerActive, _ := beego.AppConfig.Int("redis_proxy2layer_active")
	redisProxy2LayerIdleTimeout, _ := beego.AppConfig.Int("redis_proxy2layer_idle_timeout")



	logs.Debug("read config succ, redis addr :%v", redisProxy2LayerAdrr)


	secKillConf.RedisProxy2LayerConf.RedisAddr = redisBlackAdrr
	secKillConf.RedisProxy2LayerConf.RedisMaxIdle = redisProxy2LayerIdle
	secKillConf.RedisProxy2LayerConf.RedisMaxActive = redisProxy2LayerActive
	secKillConf.RedisProxy2LayerConf.RedisIdleTimeout = redisProxy2LayerIdleTimeout


	writeGoNums, err := beego.AppConfig.Int("write_proxy2layer_goroutine_num")
	if err != nil {
		logs.Debug("read writegonums err :%v", err)
		return
	}
	secKillConf.WriteProxy2LayerGoroutine_num = writeGoNums
	readGoNums, err := beego.AppConfig.Int("write_proxy2layer_goroutine_num")
	if err != nil {
		logs.Debug("read readgonums err :%v", err)
		return
	}
	secKillConf.ReadProxy2LayerGoroutine_num = readGoNums

	return

}
