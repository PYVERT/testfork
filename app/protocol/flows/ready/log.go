package ready

import (
	"github.com/karlsend/PYVERT/testfork/karlsend/infrastructure/logger"
	"github.com/karlsend/PYVERT/testfork/karlsend/util/panics"
)

var log = logger.RegisterSubSystem("PROT")
var spawn = panics.GoroutineWrapperFunc(log)

