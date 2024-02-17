package main

import (
	"github.com/karlsend/PYVERT/testfork/karlsend/infrastructure/logger"
	"github.com/karlsend/PYVERT/testfork/karlsend/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("RPIC")
	spawn      = panics.GoroutineWrapperFunc(log)
)

