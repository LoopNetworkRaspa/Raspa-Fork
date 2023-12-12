package prefixmanager

import (
	"github.com/LoopNetworkRaspa/Raspa-Fork/infrastructure/logger"
	"github.com/LoopNetworkRaspa/Raspa-Fork/util/panics"
)

var log = logger.RegisterSubSystem("PRFX")
var spawn = panics.GoroutineWrapperFunc(log)
