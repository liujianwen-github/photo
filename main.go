package main

import (
	. "./router"
)

func main() {

	router := InitRouter()
	router.Run("localhost:8000")
	println("server is running at 8000")
	//初始化congtroller
	//controller.InitController()
}
