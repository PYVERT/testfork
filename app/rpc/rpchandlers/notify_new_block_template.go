package rpchandlers

import (
	"github.com/karlsend/PYVERT/testfork/karlsend/app/appmessage"
	"github.com/karlsend/PYVERT/testfork/karlsend/app/rpc/rpccontext"
	"github.com/karlsend/PYVERT/testfork/karlsend/infrastructure/network/netadapter/router"
)

// HandleNotifyNewBlockTemplate handles the respectively named RPC command
func HandleNotifyNewBlockTemplate(context *rpccontext.Context, router *router.Router, _ appmessage.Message) (appmessage.Message, error) {
	listener, err := context.NotificationManager.Listener(router)
	if err != nil {
		return nil, err
	}
	listener.PropagateNewBlockTemplateNotifications()

	response := appmessage.NewNotifyNewBlockTemplateResponseMessage()
	return response, nil
}

