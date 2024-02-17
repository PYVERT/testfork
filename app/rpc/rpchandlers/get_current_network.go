package rpchandlers

import (
	"github.com/karlsend/PYVERT/testfork/karlsend/app/appmessage"
	"github.com/karlsend/PYVERT/testfork/karlsend/app/rpc/rpccontext"
	"github.com/karlsend/PYVERT/testfork/karlsend/infrastructure/network/netadapter/router"
)

// HandleGetCurrentNetwork handles the respectively named RPC command
func HandleGetCurrentNetwork(context *rpccontext.Context, _ *router.Router, _ appmessage.Message) (appmessage.Message, error) {
	response := appmessage.NewGetCurrentNetworkResponseMessage(context.Config.ActiveNetParams.Net.String())
	return response, nil
}

