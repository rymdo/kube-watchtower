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

	logger.Info("deployments")
	for _, item := range k.GetDeployments() {
		logger.Infof("%+v", item)
	}

	logger.Info("stagefulsets")
	for _, item := range k.GetStatefulsets() {
		logger.Infof("%+v", item)
	}

	logger.Info("daemonsets")
	for _, item := range k.GetDaemonsets() {
		logger.Infof("%+v", item)
	}
}
