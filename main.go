package main

import(
	. "./router"
)

func main()  {

	router := InitRouter()
	router.Run("localhost:8080")
	println("server is running at 8000")
	//初始化congtroller
	//controller.InitController()
}