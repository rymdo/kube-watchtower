package utils

import (
	"context"
	"fmt"
	"os"
)

func NewConfig(logger *Logger) *Config {
	config := Config{}

	// Context
	config.Ctx = context.Background()

	// Kubernetes
	config.KubeconfigPath = os.Getenv("KUBECONFIG")
	if config.KubeconfigPath == "" {
		config.KubeconfigPath = fmt.Sprintf("%s/.kube/config", getHomeDir())
	}
	if _, err := os.Stat(config.KubeconfigPath); err == nil {
		config.KubeconfigExists = true
	} else {
		config.KubeconfigExists = false
	}
	logger.Infof("%+v", config)
	return &config
}

func getHomeDir() string {
	dir, err := os.UserHomeDir()
	if err != nil {
		panic(err.Error())
	}
	return dir
}
