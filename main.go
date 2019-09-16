package main

import (
	"Graphy/info"
	"Graphy/test/grid"
	"runtime"
)

var points []info.Vec2

const ()

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func main() {

	grid.GridMain()

}
