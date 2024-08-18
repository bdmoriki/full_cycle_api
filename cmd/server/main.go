package main

import "github.com/bdmoriki/full_cycle_api/configs"

func main() {
	config, err := configs.LoadConfig(".env")
	if err != nil {
		panic(err)
	}
}
