package main

import (
	"fmt"

	"github.com/flp-fernandes/9-APIS/configs"
)

func main() {
	config, _ := configs.LoadConfig(".")
	fmt.Println(config)
}
