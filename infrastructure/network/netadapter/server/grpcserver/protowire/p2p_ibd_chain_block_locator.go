package protowire

import (
	"github.com/karlsend/PYVERT/testfork/karlsend/app/appmessage"
	"github.com/pkg/errors"
)

func (x *KarlsendMessage_IbdChainBlockLocator) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "KarlsendMessage_IbdChainBlockLocator is nil")
	}
	return x.IbdChainBlockLocator.toAppMessage()
}

func (x *IbdChainBlockLocatorMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "IbdChainBlockLocatorMessage is nil")
	}
	blockLocatorHashes, err := protoHashesToDomain(x.BlockLocatorHashes)
	if err != nil {
		return nil, err
	}
	return &appmessage.MsgIBDChainBlockLocator{
		BlockLocatorHashes: blockLocatorHashes,
	}, nil
}

func (x *KarlsendMessage_IbdChainBlockLocator) fromAppMessage(message *appmessage.MsgIBDChainBlockLocator) error {
	x.IbdChainBlockLocator = &IbdChainBlockLocatorMessage{
		BlockLocatorHashes: domainHashesToProto(message.BlockLocatorHashes),
	}
	return nil
}

