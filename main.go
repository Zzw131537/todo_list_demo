package main

import (
	"fmt"
	"todo_list_demo/config"
	"todo_list_demo/handler"
)

func main() {
	config.Init()

	fmt.Println("初始化完成")

	router := handler.NewRouter()
	router.Run(":3001")

}
