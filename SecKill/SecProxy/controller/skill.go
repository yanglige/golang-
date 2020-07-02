package controller

import (
	"SecKill/SecProxy/service"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"strconv"
	"strings"
	"time"
)

type SkillController struct {
	beego.Controller
}

func (s *SkillController) SecKill() {
	productId, err := s.GetInt("product_id")
	result := make(map[string]interface{})

	result["code"] = 0
	result["message"] = "success"
	defer func() {
		s.Data["json"] = result
		s.ServeJSON()
	}()

	if err != nil {
		result["code"] = 1001
		result["message"] = "invalid product_id"
		return
	}
	source := s.GetString("src")
	authcode := s.GetString("authcode")
	secTime := s.GetString("time")
	nance := s.GetString("nance")

	secRequest := &service.SecRequest{}
	secRequest.AuthCode = authcode
	secRequest.Nance = nance
	secRequest.ProductId = productId
	secRequest.SecTime = secTime
	secRequest.Source = source
	secRequest.UserAuthSign = s.Ctx.GetCookie("userAuthSign")
	secRequest.UserId, err = strconv.Atoi(s.Ctx.GetCookie("userId"))
	secRequest.AccessTime = time.Now()
	if len(s.Ctx.Request.RemoteAddr) > 0 {
		secRequest.ClientAddr =strings.Split(s.Ctx.Request.RemoteAddr,":")[0]
	}
	secRequest.ClientRefence = s.Ctx.Request.Referer()

	logs.Debug("client request：[%v]", secRequest)
	if err != nil {
		result["code"] = service.ErrInvalidRequest
		result["message"] = fmt.Sprintf("invalid cookie:userId")
		return
	}
	data, code, err := service.SecKill(secRequest)
	if err != nil {
		result["code"] = code
		result["message"] = err.Error()
		return
	}
	result["data"] = data
	result["code"] = code
	return

}
func (s *SkillController) SecInfo() {
	productId, err := s.GetInt("product_id")
	result := make(map[string]interface{})
	result["code"] = 0
	result["message"] = "success"
	defer func() {
		s.Data["json"] = result
		s.ServeJSON()
	}()
	if err != nil {
		data, code, err := service.SecInfoList()
		if err != nil {
			result["code"] = code
			result["message"] = err.Error()
			logs.Error("invalid reques, get product_id failed, err:%v", err)
			return
		}
		result["code"] = code
		result["data"] = data

	}else {
		data, code, err := service.SecInfo(productId)
		if err != nil {
			result["code"] = code
			result["message"] = err.Error()
			logs.Error("invalid request, get product__id failed, err:%v")
			return
		}
		result["code"] = code
		result["data"] = data

	}

}