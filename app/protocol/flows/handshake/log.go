package handshake

import (
	"github.com/LoopNetworkRaspa/Raspa-Fork/infrastructure/logger"
	"github.com/LoopNetworkRaspa/Raspa-Fork/util/panics"
)

var log = logger.RegisterSubSystem("PROT")
var spawn = panics.GoroutineWrapperFunc(log)
