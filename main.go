package main

import (
	"context"
	"fmt"
	"todo_list_demo/config"
	"todo_list_demo/handler"
	"todo_list_demo/service"
)

func main() {
	config.Init()
	ctx := context.Background()
	go service.Notice(ctx)
	fmt.Println("初始化完成")
	router := handler.NewRouter()
	router.Run(":3001")

}
