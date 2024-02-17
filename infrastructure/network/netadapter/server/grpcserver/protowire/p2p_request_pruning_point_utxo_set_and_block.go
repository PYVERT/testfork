package protowire

import (
	"github.com/karlsend/PYVERT/testfork/karlsend/app/appmessage"
	"github.com/pkg/errors"
)

func (x *KarlsendMessage_RequestPruningPointUTXOSet) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "KarlsendMessage_RequestPruningPointUTXOSet is nil")
	}
	return x.RequestPruningPointUTXOSet.toAppMessage()
}

func (x *RequestPruningPointUTXOSetMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "RequestPruningPointUTXOSetMessage is nil")
	}
	pruningPointHash, err := x.PruningPointHash.toDomain()
	if err != nil {
		return nil, err
	}
	return &appmessage.MsgRequestPruningPointUTXOSet{PruningPointHash: pruningPointHash}, nil
}

func (x *KarlsendMessage_RequestPruningPointUTXOSet) fromAppMessage(
	msgRequestPruningPointUTXOSet *appmessage.MsgRequestPruningPointUTXOSet) error {

	x.RequestPruningPointUTXOSet = &RequestPruningPointUTXOSetMessage{}
	x.RequestPruningPointUTXOSet.PruningPointHash = domainHashToProto(msgRequestPruningPointUTXOSet.PruningPointHash)
	return nil
}

