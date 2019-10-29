package main

import (
	"Graphy/frontend/graph"
	"Graphy/utils"
	"fmt"
	"runtime"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func main() {

	//grid.GridMain()

	data, err := utils.ReadTensorFromText("[1, 2]", 1, 2)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", data)

	graph.GraphMain()
}
