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

	// TODO: Read tensor data and provide it to Main

	tensor, err := utils.ReadTensorFromFile("linearData.tensor", 42)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", tensor.Data())
	fmt.Printf("%v\n", tensor)

	graph.Main(tensor.Data().([]float32))
}
