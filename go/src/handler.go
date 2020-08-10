package main

import (
	sdk "github.com/digitalcompanion-keti/dcf-watcher/go/pb"
	_mesh "github.com/digitalcompanion-keti/dcf-watcher/go/mesh"
)

func Handler(req sdk.Request) string {

	// mesh call
	// 
	// functionName := "<FUNCTIONNAME>"
	// input := string(req.input)
	// result := mesh.MeshCall(functionName, []byte(input))
	// return result


	// single call
	return string(req.input)
}
