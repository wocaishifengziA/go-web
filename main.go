package main

import (
	"fmt"
	"time"
)

func main() {
	// var m sync.Map
	// m.Store("a", 1)
	// m.Store("b", 2)
	// m.Store("c", 3)
	// fmt.Println(m.Load("a"))

	// mx := make(map[string]int)
	// m.Range(func(k, v interface{}) bool {
	// 	mx[k.(string)] = v.(int)
	// 	return true
	// })
	// fmt.Println(mx)

	// r := gin.Default()
	// g := r.Group("/cnm")
	// g.GET("/bb", func(ctx *gin.Context) {
	// 	ctx.JSON(200, "sdsd")
	// })
	// r.GET("/download", func(c *gin.Context) {
	// 	c.File("./README.pdf")
	// })
	// r.Run(":8080")

	// pflag test
	// pflagx.PlagDo()
	fmt.Println(time.Now().UnixNano()/1e6)
}
