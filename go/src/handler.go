package main

import (
	// mesh "github.com/digitalcompanion-keti/dcf-watcher/go/mesh"
	sdk "github.com/digitalcompanion-keti/dcf-watcher/go/pb"
)

func Handler(req sdk.Request) string {

	// mesh call
	// When using, uncomment the mesh that is commented out on the import.
	//
	// functionName := "<FUNCTIONNAME>"
	// input := string(req.Input)
	// result := mesh.MeshCall(functionName, []byte(input))
	// return result

	// single call
	return string(req.Input)
}
