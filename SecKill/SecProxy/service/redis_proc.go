package service

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
)

func WriteHandle() {

	for {
		req := <- secSkillConf.SecReqChan
		conn := secSkillConf.proxy2LayerRedisPool.Get()
		data, err := json.Marshal(req)
		if err != nil {
			logs.Error("json.marshal failed, err:%v, req :%v", err, req)
			continue
		}

		_, err = conn.Do("LPUSH", "sec_queue", data)
		if err != nil {
			logs.Error("lpush failed, err:%v, req:%v", err, req)
			continue
		}
		conn.Close()
	}

}


func ReadHandle() {
	return
}
