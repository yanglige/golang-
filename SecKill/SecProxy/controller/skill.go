package controller

import "github.com/astaxie/beego"

type SkillController struct {
	beego.Controller
}

func (s *SkillController) SecKill() {
	s.Data["json"] = "sec kill"
	s.ServeJSON()
}
func (s *SkillController) SecInfo() {
	s.Data["json"] = "sec info"
	s.ServeJSON()
}