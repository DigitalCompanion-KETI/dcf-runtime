package main

import sdk "github.com/digitalcompanion-keti/dcf-watcher/go/pb"

func Handler(req sdk.Request) string {
	return string(req.Input)
}
