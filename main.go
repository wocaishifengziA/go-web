package main

import (
	"fmt"
)

type Toy interface {
	Play()
}

type Puzzle struct {
	Name string
}

func (*Puzzle) Play() {
	fmt.Println("palying puzzle!")
}

func main() {
	// configs.InitConfig("./conf/config.yaml")
	// loggers.InitLogger(configs.Config.Log)
	// loggers.LogInstance().Infoln("ok")
	// rabbitmq.DoBlock()
	// r := gin.Default()
	// r := gin.New()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	// syncpool.Run()
	var t Toy
	p := &Puzzle{Name: "jake"}
	t = p
	t.Play()
	x, ok := t.(*Puzzle)
	if !ok {
		fmt.Println("not ok")
	}
	fmt.Println("xx: ", x.Name)
}
