package config

import (
	"context"
	"fmt"
	"os"
)

func init() {
	config = &Config{}

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
}

func getHomeDir() string {
	dir, err := os.UserHomeDir()
	if err != nil {
		panic(err.Error())
	}
	return dir
}

func GetConfig() *Config {
	return config
}
