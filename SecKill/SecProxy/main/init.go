package main

import (
	"SecKill/SecProxy/service"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	etcd_client "github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"github.com/gomodule/redigo/redis"
	"golang.org/x/net/context"
	"time"
)

var (
	redisPool *redis.Pool
	etcdClient *etcd_client.Client


)
func initRedis() (err error){
	redisPool = &redis.Pool{
		MaxIdle: secKillConf.RedisConf.RedisMaxIdle,
		MaxActive: secKillConf.RedisConf.RedisMaxActive,
		IdleTimeout: time.Duration(secKillConf.RedisConf.RedisIdleTimeout)*time.Second,
		Dial: func() ( redis.Conn,  error) {
			return redis.Dial("tcp", secKillConf.RedisConf.RedisAddr)
		},
	}
	conn := redisPool.Get()
	defer conn.Close()
	_, err = conn.Do("ping")
	if err != nil {
		logs.Error("ping redis failed, %v", err)
	}

	return
}

func initEtcd() (err error) {
	cli, err := etcd_client.New(etcd_client.Config{
		Endpoints: []string{secKillConf.EtcdConf.EtcdAddr},
		DialTimeout: time.Duration(secKillConf.EtcdConf.Timeout) * time.Second,
	})
	if err != nil {
		logs.Error("connect etcd failed, err:", err)
		return
	}
	etcdClient = cli
	return
}

func convertLogLevel(level string) int {
	switch level {
	case "debug":
		return logs.LevelDebug
	case "warn":
		return logs.LevelWarn
	case "info":
		return logs.LevelInfo
	case "trace":
		return logs.LevelTrace
	}
	return logs.LevelDebug
}
func initLogger() (err error) {
	fmt.Println("begin initLogger")
	config := make(map[string]interface{})
	config["filename"] = secKillConf.LogPath
	config["level"] = convertLogLevel(secKillConf.LogLevel)

	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println("marshal failed, err:", err)
		return
	}
	logs.SetLogger(logs.AdapterFile, string(configStr))
	//fmt.Println("finish initlogger")
	return
}

func loadSecConf() (err error) {
	fmt.Println("begin loadsecconf")
	//key := fmt.Sprintf("%s/product", secKillConf.etcdConf.etcdSecKey)
	resp, err := etcdClient.Get(context.Background(), secKillConf.EtcdConf.EtcdSecProductKey)
	if err != nil {
		logs.Error("get [%s] from etcd failed, err :%v", secKillConf.EtcdConf.EtcdSecProductKey, err)
		return
	}
	fmt.Println("---------------")
	var secProductInfo []service.SecProductInfoConf

	for k,v := range resp.Kvs {
		logs.Debug("key[%s] valued[%v]", k, v)
		err = json.Unmarshal(v.Value, &secProductInfo)
		if err != nil {
			logs.Error("unmarshal sec product info failed, err:%v", err)
			return
		}
		logs.Debug("sec info conf is [%v]", secProductInfo)
	}
	fmt.Println("+++++++++++++")
	//updateSecProductInfo(secProductInfo)

	secKillConf.RWSecProductLock.Lock()
	for _, v := range secProductInfo {
		secKillConf.SecProductInfoMap[v.ProductId] = &v
	}
	secKillConf.RWSecProductLock.Unlock()
	return
}
func initSec() (err error) {

	err = initLogger()
	if err != nil {
		logs.Error("init logger failed, err:%v", err)
	}
	fmt.Println("initLogger succ")
	//err = initRedis()
	//if err != nil {
	//	logs.Error("init redis failed, err:%v", err)
	//	return
	//}

	err = initEtcd()
	if err != nil {
		logs.Error("init etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("initEtcd succ")

	err = loadSecConf()
	if err != nil {
		logs.Error("load sec conf failed err:%v\n", err)
		return
	}
	fmt.Println("initSecconf succ")

	service.InitService(secKillConf)
	initSecProductWatch()

	logs.Info("init sec succ")

	return
}


func initSecProductWatch() {
	go watchSecProductKey(secKillConf.EtcdConf.EtcdSecProductKey)

}

func watchSecProductKey(key string) {
	cli, err := etcd_client.New(etcd_client.Config{
		Endpoints:[]string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout:5 * time.Second,
	})
	if err != nil {
		logs.Error("connect etcd failed, err:", err)
		return
	}
	logs.Debug("begin watch key:%s", key)
	for {
		rch := cli.Watch(context.Background(), key)
		var secProductInfo []service.SecProductInfoConf
		var getConfSucc = true
		for wresp := range rch {
			for _, ev := range wresp.Events {
				if ev.Type == mvccpb.DELETE {
					logs.Warn("key[%s] is config deleted", key)
					continue
				}
				if ev.Type == mvccpb.PUT && string(ev.Kv.Key) == key {
					err = json.Unmarshal(ev.Kv.Value, &secProductInfo)
					if err != nil {
						logs.Error("key [%s], unmarshal[%s], err:%v", err)
						getConfSucc = false
						continue
					}
				}
				logs.Debug("get config from etcd, %s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)

			}
			if getConfSucc {
				logs.Debug("get config from etcd succ, %v", secProductInfo)
				updateSecProductInfo(secProductInfo)
			}
		}
	}
}
func updateSecProductInfo(secProductInfo []service.SecProductInfoConf) {

	// 优化加锁效率低的方法：
	var tmp map[int]*service.SecProductInfoConf = make(map[int]*service.SecProductInfoConf, 1024)
	for _, v := range secProductInfo {
		tmp[v.ProductId] = &v
	}
	secKillConf.RWSecProductLock.Lock()
	secKillConf.SecProductInfoMap = tmp
	secKillConf.RWSecProductLock.Unlock()


}


