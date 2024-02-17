package main

import "github.com/karlsend/PYVERT/testfork/karlsend/cmd/karlsenwallet/daemon/server"

func startDaemon(conf *startDaemonConfig) error {
	return server.Start(conf.NetParams(), conf.Listen, conf.RPCServer, conf.KeysFile, conf.Profile, conf.Timeout)
}

