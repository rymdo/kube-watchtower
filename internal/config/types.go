package config

import "context"

var (
	config *Config
)

type Config struct {
	Ctx              context.Context
	KubeconfigPath   string
	KubeconfigExists bool
}
