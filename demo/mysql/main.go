package main

import "github.com/wocaishifengziA/go-web/demo/mysql/database"

func main() {
	err := database.DbConnect()
	if err != nil {
		panic(err)
	}
}
