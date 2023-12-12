package rejects

import (
	"github.com/LoopNetworkRaspa/Raspa-Fork/app/appmessage"
	"github.com/LoopNetworkRaspa/Raspa-Fork/app/protocol/protocolerrors"
	"github.com/LoopNetworkRaspa/Raspa-Fork/infrastructure/network/netadapter/router"
)

// HandleRejectsContext is the interface for the context needed for the HandleRejects flow.
type HandleRejectsContext interface {
}

type handleRejectsFlow struct {
	HandleRejectsContext
	incomingRoute, outgoingRoute *router.Route
}

// HandleRejects handles all reject messages coming through incomingRoute.
// This function assumes that incomingRoute will only return MsgReject.
func HandleRejects(context HandleRejectsContext, incomingRoute *router.Route, outgoingRoute *router.Route) error {
	flow := &handleRejectsFlow{
		HandleRejectsContext: context,
		incomingRoute:        incomingRoute,
		outgoingRoute:        outgoingRoute,
	}
	return flow.start()
}

func (flow *handleRejectsFlow) start() error {
	message, err := flow.incomingRoute.Dequeue()
	if err != nil {
		return err
	}
	rejectMessage := message.(*appmessage.MsgReject)

	return protocolerrors.Errorf(false, "got reject message: `%s`", rejectMessage.Reason)
}
