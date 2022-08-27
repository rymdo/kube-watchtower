package utils

import (
	"context"

	"go.uber.org/zap"
)

type Config struct {
	Ctx              context.Context
	KubeconfigPath   string
	KubeconfigExists bool
}

type Logger struct {
	*zap.SugaredLogger
}
