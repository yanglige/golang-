package service

import (
	"github.com/astaxie/beego/logs"
	"github.com/gomodule/redigo/redis"
	"strconv"
	"time"
)

func InitService(serviceConf *SecSkillConf) (err error) {

	secSkillConf = serviceConf
	err = loadBlackList()
	if err != nil {
		logs.Error("load blacklist err:%v", err)
		return
	}
	logs.Debug("init service succ, config:%v", secSkillConf)
	err = initProxy2LayerRedis()
	if err != nil {
		logs.Error("load initProxy2LayerRedis err:%v", err)
		return
	}

	initRedisProcessFunc()
	secSkillConf.SecReqChan = make(chan *SecRequest, secSkillConf.SecReqChanSize)

	return
	//secSkillConf.secLimitMgr = &secLimitMgr{
	//	UserLimitMap : make(map[int]*Limit, )
	//
	//}

}
func initRedisProcessFunc() {
	for i := 0; i < secSkillConf.WriteProxy2LayerGoroutine_num; i++ {
		go WriteHandle()
	}
	for i := 0; i < secSkillConf.ReadProxy2LayerGoroutine_num; i++ {
		go ReadHandle()
	}
}

func initBlackRedis() (err error){
	secSkillConf.blackRedisPool = &redis.Pool{
		MaxIdle: secSkillConf.RedisBlackConf.RedisMaxIdle,
		MaxActive: secSkillConf.RedisBlackConf.RedisMaxActive,
		IdleTimeout: time.Duration(secSkillConf.RedisBlackConf.RedisIdleTimeout)*time.Second,
		Dial: func() ( redis.Conn,  error) {
			return redis.Dial("tcp", secSkillConf.RedisBlackConf.RedisAddr)
		},
	}
	conn := secSkillConf.blackRedisPool.Get()
	defer conn.Close()
	_, err = conn.Do("ping")
	if err != nil {
		logs.Error("ping redis failed, %v", err)
	}

	return
}
func initProxy2LayerRedis() (err error) {
	secSkillConf.proxy2LayerRedisPool = &redis.Pool{
		MaxIdle: secSkillConf.RedisProxy2LayerConf.RedisMaxIdle,
		MaxActive: secSkillConf.RedisProxy2LayerConf.RedisMaxActive,
		IdleTimeout: time.Duration(secSkillConf.RedisProxy2LayerConf.RedisIdleTimeout)*time.Second,
		Dial: func() ( redis.Conn,  error) {
			return redis.Dial("tcp", secSkillConf.RedisProxy2LayerConf.RedisAddr)
		},
	}
	conn := secSkillConf.proxy2LayerRedisPool.Get()
	defer conn.Close()
	_, err = conn.Do("ping")
	if err != nil {
		logs.Error("ping redis failed, %v", err)
	}

	return
}
func loadBlackList() (err error) {
	secSkillConf.ipBlackMap = make(map[string]bool, 10000)
	secSkillConf.idBlackMap = make(map[int]bool, 10000)
	err = initBlackRedis()
	if err != nil {
		logs.Error("initBlackREdis err :%v", err)
		return
	}
	conn := secSkillConf.blackRedisPool.Get()
	defer conn.Close()
	reply, err := conn.Do("hgetall", "idblacklist")
	idlist, err := redis.Strings(reply, err)
	if err != nil {
		logs.Warn("hget all failed, err:%v", err)
		return
	}
	for _, v := range idlist {
		id, err := strconv.Atoi(v)
		if err != nil {
			logs.Warn("invalid user id [%v]", id)
			continue
		}
		secSkillConf.idBlackMap[id] = true
	}
	reply, err = conn.Do("hgetall", "ipblacklist")
	iplist, err := redis.Strings(reply, err)
	if err != nil {
		logs.Warn("hget all failed, err:%v", err)
		return
	}
	for _, v := range iplist {
		secSkillConf.ipBlackMap[v] = true
	}
	go SyncIpBlackList()
	go SyncIdBlackList()
	return
}
func SyncIpBlackList() {
	var ipList []string
	lastTime := time.Now().Unix()
	for {
		conn := secSkillConf.blackRedisPool.Get()
		defer conn.Close()
		reply, err := conn.Do("BLPOP", "blackiplist", time.Second)
		ip, err := redis.String(reply, err)
		if err != nil {
			continue
		}
		curTime := time.Now().Unix()
		ipList = append(ipList, ip)
		if len(ipList) > 100 || curTime - lastTime > 5 {
			secSkillConf.RWBlackLock.Lock()
			for _, v := range ipList {
				secSkillConf.ipBlackMap[v] = true
			}
			secSkillConf.RWBlackLock.Unlock()
			lastTime = curTime
			logs.Info("sync ip list from redis succ, ip[%v]", ipList)
		}

	}
}
func SyncIdBlackList() {
	for {
		conn := secSkillConf.blackRedisPool.Get()
		defer conn.Close()
		reply, err := conn.Do("BLPOP", "blackidlist", time.Second)
		id, err := redis.Int(reply, err)
		if err != nil {
			continue
		}
		secSkillConf.RWBlackLock.Lock()
		secSkillConf.idBlackMap[id] = true
		secSkillConf.RWBlackLock.Unlock()
	}
}