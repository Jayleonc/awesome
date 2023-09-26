package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	r := gin.Default()
	r.Use(gsgsfhfa)
	// 每当有一个请求到达服务端时, Gin就会为这个请求分配一个Context, 该上下文中保存了这个请求对应的处理器链, 以及一个索引idnex用于记录当前处理到了哪个HandlerFunc, index的初始值为-1
	group := r.Group("/v1").Use(groupUse).Use(groupUse2)
	{
		group.Use(groupUse3).GET("/g1", funcName())
		group.GET("/g2", funcName2())
		group.GET("/g3", funcName())
		group.GET("/g4", funcName())
		group.GET("/g5", funcName())
		group.GET("/g6", funcName())
		group.GET("/g7", funcName())
		group.GET("/g8", funcName())
		group.GET("/g9", funcName())
	}

	group2 := r.Group("/v2").Use(groupUse)
	{
		group2.GET("/g1", funcName())
		group2.GET("/g2", funcName())
		group2.GET("/g3", funcName())
		group2.GET("/g4", funcName())
		group2.GET("/g5", funcName())
		group2.GET("/g6", funcName())
		group2.GET("/g7", funcName())
		group2.GET("/g8", funcName())
		group2.GET("/g9", funcName())
	}

	r.Run()
}

func gsgsfhfa(context *gin.Context) {
	start := time.Now()
	fmt.Println("time first")
	context.Next()
	fmt.Printf("time cost %d ms\n", time.Since(start).Milliseconds())
}
func groupUse(c *gin.Context) {
	fmt.Printf("before handle\n")
	c.Next()
	fmt.Printf("two handle\n")
	c.Next()
	fmt.Printf("after handle\n")
}
func groupUse2(c *gin.Context) {
	fmt.Printf("22222222222\n")
	c.Next()
}
func groupUse3(c *gin.Context) {
	fmt.Printf("3333333333\n")
	c.Next()
}

func funcName() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"me": "asd",
		})
	}
}
func funcName2() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"me": "asd",
		})
	}
}
