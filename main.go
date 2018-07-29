package main

import (
	. "./router"
)

func main() {

	router := InitRouter()
	router.Run(":8001")
	println("server is running at 8001")
	//初始化congtroller
	//controller.InitController()
}
