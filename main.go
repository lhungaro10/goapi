package main

import (
	"github.com/lhungaro10/goapi/config"
	"github.com/lhungaro10/goapi/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// Initialize configs
	configError := config.Init();

	if configError != nil {
		logger.ErrorF("failed to initialize config: %v", configError )
		return
	}

	router.Initialize()
}
