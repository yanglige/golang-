package controller

import (
	"SecKill/SecProxy/service"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type SkillController struct {
	beego.Controller
}

func (s *SkillController) SecKill() {
	s.Data["json"] = "sec kill"
	s.ServeJSON()
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

		result["code"] = 1001
		result["messega"] = "invalied product_id"


		logs.Error("invalid request, get product_id failed, err:%v")
		return
	}
	data, code, err := service.SecInfo(productId)
	if err != nil {
		result["code"] = code
		result["message"] = err.Error()
		logs.Error("invalid request, get product__id failed, err:%v")
		return
	}

	result["data"] = data

}