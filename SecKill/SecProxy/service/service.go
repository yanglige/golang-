package service

import (
	"crypto/md5"
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

func SecInfoList() (data []map[string]interface{}, code int, err error) {
	secSkillConf.RWSecProductLock.RLock()
	defer secSkillConf.RWSecProductLock.RUnlock()
	for _, v := range secSkillConf.SecProductInfoMap {
		logs.Debug("get product[%d]",v.ProductId)
		item, _, err := SecInfoById(v.ProductId)
		if err != nil {
			logs.Error("get product_id[%d] failed, err:%v", v.ProductId, err)
			continue
		}
		logs.Debug("get product[%d], result[%v], all[%v], v[%v]", v.ProductId, item, secSkillConf.SecProductInfoMap, v)
		data = append(data, item)
	}
	return
}

func SecInfo(productId int) (data []map[string]interface{}, code int, err error) {
	secSkillConf.RWSecProductLock.RLock()
	defer secSkillConf.RWSecProductLock.RUnlock()

	item, code, err := SecInfoById(productId)
	if err != nil {
		return
	}
	data = append(data, item)
	return

}


func SecKill(req *SecRequest) (data []map[string]interface{}, code int, err error) {
	secSkillConf.RWSecProductLock.RLock()
	defer secSkillConf.RWSecProductLock.RUnlock()
    // 用户校验
	err = userCheck(req)
	if err != nil {
		code = ErrUserCheckAuthFailed
		logs.Warn("userId[%d] invalid, check failed, req[%v]", req.UserId, req)
		return
	}
	// 频率控制
	err = antiSpam(req)
	if err != nil {
		code = ErrUserCheckAuthFailed
		logs.Warn("userId[%d] invalid, userserviceBusy, req[%v]", req.UserId, req)
		return
	}

	return
}


func userCheck(req *SecRequest) (err error) {
	found := false
	for _, refer := range secSkillConf.ReferWhiteList {
		if refer == req.ClientRefence {
			found = true
			break
		}
	}
	if !found {
		err = fmt.Errorf("invalid request")
		logs.Warn("user[%d] is reject by refer, req[%v]", req.UserId, req)
	    return
	}



	authData := fmt.Sprintf("%d:%s", req.UserId,secSkillConf.CookieSecretKey)
	authSign := fmt.Sprintf("%x", md5.Sum([]byte(authData)))
	if authSign != req.UserAuthSign {
		err = fmt.Errorf("invalid user cookie auth")
		return
	}
	return
}
func SecInfoById(productId int) (data map[string]interface{}, code int, err error) {
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