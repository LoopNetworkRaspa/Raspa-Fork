package connmanager

import (
	"github.com/LoopNetworkRaspa/Raspa-Fork/infrastructure/logger"
	"github.com/LoopNetworkRaspa/Raspa-Fork/util/panics"
)

var log = logger.RegisterSubSystem("CMGR")
var spawn = panics.GoroutineWrapperFunc(log)
