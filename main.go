package main

import (
	"beeproject/controllers"
	"github.com/astaxie/beego"
)

func main() {
	beego.BeeApp.RegisterController("/", &controllers.MainController{})
	beego.BeeApp.Run()
}

