package main

import (
	"github.com/astaxie/beego"
	_ "Todo-Go/initial"
	_ "Todo-Go/routers"
)

func main() {
	beego.Run()
}

