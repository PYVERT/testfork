package model

import "github.com/karlsend/PYVERT/testfork/karlsend/domain/consensus/model/externalapi"

// FinalityManager provides method to validate that a block does not violate finality
type FinalityManager interface {
	VirtualFinalityPoint(stagingArea *StagingArea) (*externalapi.DomainHash, error)
	FinalityPoint(stagingArea *StagingArea, blockHash *externalapi.DomainHash, isBlockWithTrustedData bool) (*externalapi.DomainHash, error)
}

