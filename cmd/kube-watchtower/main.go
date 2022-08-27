package main

import (
	"github.com/rymdo/kube-watchtower/v2/internal/kubernetes"
	"github.com/rymdo/kube-watchtower/v2/internal/utils"
)

func main() {
	logger := utils.NewLogger()
	logger.Info("started")
	cfg := utils.NewConfig(logger)
	k := kubernetes.New(logger, cfg)
	k.GetNodes()
}
