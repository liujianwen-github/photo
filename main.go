package main

import (
	. "./router"
	"fmt"
)

func main() {

	router := InitRouter()
	router.Run(":8001")
	println("server is running at 8001")
	//初始化congtroller
	//controller.InitController()
}
