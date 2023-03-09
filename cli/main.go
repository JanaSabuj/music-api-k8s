package main

import (
	"fmt"
	"github.com/JanaSabuj/music-api-k8s/pkg/config"
)

func main() {
	fmt.Println("Welcome to music api...")

	cfg, err := config.NewConfig()
	fmt.Println(cfg, err)
}
