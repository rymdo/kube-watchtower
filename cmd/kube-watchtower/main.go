package main

import (
	"fmt"

	"github.com/rymdo/kube-watchtower/v2/internal/config"
	"github.com/rymdo/kube-watchtower/v2/internal/kubernetes"
)

func main() {
	fmt.Println("Hello, World!")

	cfg := config.GetConfig()

	_ = kubernetes.New(cfg)
}
