package main

import "github.com/xcrossed/create-go-spring/create"

func main() {
	//初始化
	opts := create.NewOptions()
	creator := create.NewCreate(opts)
	creator.Run()
}
