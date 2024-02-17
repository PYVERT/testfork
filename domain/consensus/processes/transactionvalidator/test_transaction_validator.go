package transactionvalidator

import (
	"github.com/karlsend/PYVERT/testfork/karlsend/domain/consensus/model"
	"github.com/karlsend/PYVERT/testfork/karlsend/domain/consensus/model/testapi"
	"github.com/karlsend/PYVERT/testfork/karlsend/domain/consensus/utils/txscript"
)

type testTransactionValidator struct {
	*transactionValidator
}

// NewTestTransactionValidator creates an instance of a TestTransactionValidator
func NewTestTransactionValidator(baseTransactionValidator model.TransactionValidator) testapi.TestTransactionValidator {
	return &testTransactionValidator{transactionValidator: baseTransactionValidator.(*transactionValidator)}
}

func (tbv *testTransactionValidator) SigCache() *txscript.SigCache {
	return tbv.sigCache
}

func (tbv *testTransactionValidator) SetSigCache(sigCache *txscript.SigCache) {
	tbv.sigCache = sigCache
}

