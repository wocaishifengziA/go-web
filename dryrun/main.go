package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
)

type VideosetLoaderMeta struct {
	ImagesetID  int `json:"imagesetId"`
	ImageTaskId int `json:"imageTaskId"`
}

func main() {
	// v := VideosetLoaderMeta{
	// 	ImagesetID:  1,
	// 	ImageTaskId: 2,
	// }
	// PrintJson(v, "a")
	// var x int
	// a := math.Ceil(float64(x) / float64(10))
	// fmt.Println(a)
	// GetSlice(19, 10)
	var err error
	defer func() {
		// if err := recover(); err != nil {
		// 	fmt.Println(err)
		// }
		if err != nil {
			fmt.Println("err != nil")
		}
	}()
	a := []int{0, 1}
	err = errors.New("kkk")
	fmt.Println(err.Error())
	err = nil
	fmt.Println(a[2])
}

func PrintJson(data interface{}, msg string) {
	s, _ := json.Marshal(data)
	log.Println(msg)
	log.Println("json: ", string(s))
}

func GetSlice(total, segNum int) []int {
	log.Println("segNum:", segNum)
	log.Println("total: ", total)
	sliceArr := []int{}

	segSize := int(math.Ceil(float64(total) / float64(segNum)))
	for i := 0; i < segNum-1; i++ {
		sliceArr = append(sliceArr, segSize)
	}
	last := total - (segNum-1)*segSize
	sliceArr = append(sliceArr, last)

	log.Println("分割任务：", sliceArr)
	return sliceArr
}
