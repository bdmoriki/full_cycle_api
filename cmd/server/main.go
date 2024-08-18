package main

import "github.com/bdmoriki/full_cycle_api/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	print(config.DBDriver)
}
