package flowcontext

import (
	"github.com/karlsend/PYVERT/testfork/karlsend/domain"
)

// Domain returns the Domain object associated to the flow context.
func (f *FlowContext) Domain() domain.Domain {
	return f.domain
}

