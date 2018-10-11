package main

import sdk "github.com/digitalcompanion-keti/dcf-watcher/pb"

func Handler(req sdk.Request) string {
	return "echo:" + string(req.Input)
}
