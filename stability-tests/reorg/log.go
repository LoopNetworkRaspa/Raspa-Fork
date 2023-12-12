package main

import (
	"github.com/LoopNetworkRaspa/Raspa-Fork/infrastructure/logger"
	"github.com/LoopNetworkRaspa/Raspa-Fork/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("RORG")
	spawn      = panics.GoroutineWrapperFunc(log)
)
