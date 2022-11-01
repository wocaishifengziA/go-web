package main

import (
	"encoding/json"
	"log"
)

type VideosetLoaderMeta struct {
	ImagesetID  int `json:"imagesetId"`
	ImageTaskId int `json:"imageTaskId"`
}

func main() {
	v := VideosetLoaderMeta{
		ImagesetID:  1,
		ImageTaskId: 2,
	}
	PrintJson(v, "a")
}

func PrintJson(data interface{}, msg string) {
	s, _ := json.Marshal(data)
	log.Println(msg)
	log.Println("json: ", string(s))
}
