package main

import (
	_ "SecKill/SecProxy/router"
	"fmt"
	"github.com/astaxie/beego"
)




func main() {
	//workPath, _ := os.Getwd()
	//appConfigPath := filepath.Join(workPath, "conf", "app.conf")
	//fmt.Print(appConfigPath,"\n")
	err := initConfig()
	if err != nil {
		panic(err)
		return
	}
	fmt.Println("initconfig succ")
	err = initSec()
	if err != nil {
		panic(err)
		return
	}
	fmt.Println("initsec succ")
	beego.Run()


}