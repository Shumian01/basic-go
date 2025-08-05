package main

//func main() {
//	server := gin.Default() //创建的Engine
//	//当一个http请求 用get方法访问的时候，如果访问路径是/hello,
//	//Context等于上下文 处理输入输出
//	server.GET("/hello", func(c *gin.Context) {
//		//就执行这段代码
//		c.String(http.StatusOK, "Hello Go!")
//	})
//	//注册一个处理post请求的路由
//	server.POST("/post", func(ctx *gin.Context) {
//		ctx.String(http.StatusOK, "post方法")
//	})
//	//参数路由
//	//非restful风格 :/user/delete?name=daming
//	//restful风格
//	//get /users/daming 查询
//	//delete /users/daming 把我删掉
//	//put /users/daming 注册
//	//post/users/daming修改
//	server.GET("/users/:name", func(ctx *gin.Context) {
//		name := ctx.Param("name")
//
//		ctx.String(http.StatusOK, "Hello 这是参数路由"+name)
//	})
//	//通配符路由
//	server.GET("/views/*.html", func(ctx *gin.Context) {
//		page := ctx.Param(".html")
//		ctx.String(http.StatusOK, "Hello 这是通配符路由"+page)
//	})
//	//查询参数
//	server.GET("/order", func(ctx *gin.Context) {
//		oid := ctx.Query("id")
//		ctx.String(http.StatusOK, "Hello 查询参数"+oid)
//	})
//	server.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
//}
