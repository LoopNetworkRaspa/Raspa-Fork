package flowcontext

import (
	"github.com/LoopNetworkRaspa/Raspa-Fork/infrastructure/network/addressmanager"
)

// AddressManager returns the address manager associated to the flow context.
func (f *FlowContext) AddressManager() *addressmanager.AddressManager {
	return f.addressManager
}
