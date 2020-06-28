package router

import (
	"SecKill/SecProxy/controller"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/seckill", &controller.SkillController{}, "*:SecKill")
	beego.Router("/secinfo", &controller.SkillController{}, "*:SecInfo")

}