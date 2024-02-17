package rpchandlers

import (
	"github.com/karlsend/PYVERT/testfork/karlsend/app/appmessage"
	"github.com/karlsend/PYVERT/testfork/karlsend/app/rpc/rpccontext"
	"github.com/karlsend/PYVERT/testfork/karlsend/infrastructure/network/netadapter/router"
)

// HandleNotifyPruningPointUTXOSetOverrideRequest handles the respectively named RPC command
func HandleNotifyPruningPointUTXOSetOverrideRequest(context *rpccontext.Context, router *router.Router, _ appmessage.Message) (appmessage.Message, error) {
	listener, err := context.NotificationManager.Listener(router)
	if err != nil {
		return nil, err
	}
	listener.PropagatePruningPointUTXOSetOverrideNotifications()

	response := appmessage.NewNotifyPruningPointUTXOSetOverrideResponseMessage()
	return response, nil
}

