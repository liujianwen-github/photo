package router

import (
. "github.com/melonws/goweb/middleware"
"gopkg.in/gin-gonic/gin.v1"
"../middleWare"
. "../controller"
)

/**
 * 这个函数在main()中调用,相当于new了一个router对象,
 * 并且构造了一些属性,比如import了中间件,和apis
 * 把整体路由使用了AuthMiddleWare中间件,然后把 POST /person 指向了 AddPersonAPI这个handle函数
 * 可以在apis的person包里找到
 */
func InitRouter() *gin.Engine {

	router := gin.Default()

	//全局中间件
	router.Use(AuthMiddleWare())
	router.Use(middleWare.Cors())
	{
		router.GET("/",IndexApi)
		router.POST("/photo/item",CreateItem)
		//router.POST("/todo/item", AddTodo)
		//router.DELETE("/todo/item/:id",DeleteTodo)
		//router.PUT("/todo/item/:id",UpdateTodo)
		//router.GET("/todo/item/:id",GetTodo)
		//router.GET("/todo/list",GetList)
	}

	//群组
	//authorized := router.Group("/", AuthMiddleWare())
	//authorized.GET("/", IndexApi)

	return router
}
