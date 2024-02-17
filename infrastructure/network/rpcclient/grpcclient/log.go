package grpcclient

import (
	"github.com/karlsend/PYVERT/testfork/karlsend/infrastructure/logger"
	"github.com/karlsend/PYVERT/testfork/karlsend/util/panics"
)

var log = logger.RegisterSubSystem("RPCC")
var spawn = panics.GoroutineWrapperFunc(log)

