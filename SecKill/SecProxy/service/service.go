package service

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"time"
)

var (
	secSkillConf *SecSkillConf
)
func InitService(serviceConf *SecSkillConf) {
	secSkillConf = serviceConf
	logs.Debug("init service succ, config:%v", secSkillConf)
}


func SecInfo(productId int) (data map[string]interface{}, code int, err error) {
	secSkillConf.RWSecProductLock.RLock()
	defer secSkillConf.RWSecProductLock.RUnlock()

	v, ok := secSkillConf.SecProductInfoMap[productId]
	if !ok {
		code = ErrNotFoundProductId
		err = fmt.Errorf("not found product_id:%d", productId)
		return
	}
	start := false
	end := false
	status := "success"
	now := time.Now().Unix()
	if now - v.StartTime < 0 {
		start = false
		end = false
		status = "sec kill is not start"
	}
	if now - v.StartTime > 0 {
		start =  true
	}
	if now - v.EndTime > 0 {
		start = false
		end = true
		status = "sec kill is already end"
	}
	if v.Status == ProductStatusForceSaleOut || v.Status == ProductStatusForceSaleOut {
		start = false
		end = true
		status = "product is sale out"
	}

	data = make(map[string]interface{})
	data["product_Id"] = productId
	data["start"] = start
	data["end"] = end
	data["status"] = status
	return

}