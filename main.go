package main

import (
	"log_test/Helper"

	"go.uber.org/zap"
)

func main() {
	Helper.InitLogger()
	zap.S().Debug("initial Debug")
	zap.S().Info("initial info")
	zap.S().Warn("initial warning")
	zap.S().Error("initial error")
}
