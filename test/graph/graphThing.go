package graph

import (
	"Graphy/test"
	"Graphy/utils"
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"log"
	"runtime"
)

const (
	width  = 500
	height = 500
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func GraphMain() {

	window := initGLFW()
	program := initGL()

	points := initSinWave()

}

func initGL() uint32 {

	if err := gl.Init(); err != nil {
		log.Panicf("Unable to initialize GL %v", err)
	}

	log.Println("OpenGL version", gl.GoStr(gl.GetString(gl.VERSION)))

	fragShader, err := utils.LoadShader("shaders/graph/graph.frag")

	if err != nil {
		panic(err)
	}

	vertShader, err := utils.LoadShader("shaders/graph/graph.vert")

	if err != nil {
		panic(err)
	}

	program := gl.CreateProgram()

	gl.AttachShader(program, fragShader)
	gl.AttachShader(program, vertShader)
	gl.LinkProgram(program)

	return program
}

func initGLFW() *glfw.Window {
	// glfw.WindowHint(glfw.Resizable, glfw.True) : Default

	if err := glfw.Init(); err != nil {
		log.Panicf("Unable to initialize GLFW %v", err)
	}

	glfw.WindowHint(glfw.Visible, glfw.False)
	glfw.WindowHint(glfw.Samples, 10)
	/*
		glfw.WindowHint(glfw.ContextVersionMajor, 4)
		glfw.WindowHint(glfw.ContextVersionMinor, 1)
		glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
		glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	*/
	window, err := glfw.CreateWindow(width, height, "", nil, nil)

	if err != nil {
		log.Panicf("Unable to create GLFW %v", err)
	}

	// https://stackoverflow.com/a/3270733
	window.SetSizeCallback(func(_ *glfw.Window, width int, height int) {
		gl.Viewport(0, 0, int32(width), int32(height))
	})

	//var vertexArrayID uint32
	//gl.GenVertexArrays(1, &vertexArrayID)
	//gl.BindVertexArray(vertexArrayID)

	width, height := window.GetSize()

	vidMode := glfw.GetPrimaryMonitor().GetVideoMode()

	window.SetPos((vidMode.Width-width)/2, (vidMode.Height-height)/2)
	window.MakeContextCurrent()
	window.Show()

	glfw.SwapInterval(1) // VSync

	return window
}