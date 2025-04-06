package main

import "github.com/vinizer4/go-expert-fullcycle/apis/configs"

func main() {
	config, _ := configs.LoadConfig()
	println(config.DBDriver)
}
