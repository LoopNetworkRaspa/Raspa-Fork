// Copyright (c) 2013-2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package winservice

import (
	"github.com/LoopNetworkRaspa/Raspa-Fork/infrastructure/logger"
	"github.com/LoopNetworkRaspa/Raspa-Fork/util/panics"
)

var log = logger.RegisterSubSystem("CNFG")
var spawn = panics.GoroutineWrapperFunc(log)
