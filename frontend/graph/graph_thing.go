package graph

import (
	"Graphy/info"
	"Graphy/utils"
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"log"
	"math"
	"runtime"
	"time"
)

const (
	fps    = 60
	width  = 350
	height = 500
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func Main(points []float32) {

	window := initGLFW()
	program := initGL()

	points2 := initSineWave()

	/*vao := */
	utils.MakeVaoByVec2(points2)
	//gl.BindVertexArray(vao)

	//gl.BindVertexArray(utils.MakeVaoByVec2(points))

	//utils.NormalizeRGB(153, 211, 205) Blue thing
	r, g, b := utils.NormalizeRGB(255, 255, 255)
	gl.ClearColor(r, g, b, 1)

	for !window.ShouldClose() {

		start := time.Now()
		drawVec2(points2, program, window)
		elapsed := time.Since(start)

		//actualFPS := ((1000 / fps) * time.Millisecond) + elapsed

		time.Sleep(((1000 / fps) * time.Millisecond) - elapsed)
	}

	glfw.Terminate()
}

func drawFloat32(points []float32, program uint32, window *glfw.Window) {

	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)

	//gl.BindVertexArray(vao)
	gl.DrawArrays(gl.POINTS, 0, int32(len(points)))

	glfw.PollEvents()
	window.SwapBuffers()
}

func drawVec2(points []info.Vec2, program uint32, window *glfw.Window) {

	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)

	//gl.BindVertexArray(vao)
	gl.DrawArrays(gl.POINTS, 0, int32(len(points)))

	glfw.PollEvents()
	window.SwapBuffers()
}

func initGL() uint32 {

	if err := gl.Init(); err != nil {
		log.Panicf("Unable to initialize GL %v", err)
	}

	log.Println("OpenGL version", gl.GoStr(gl.GetString(gl.VERSION)))

	fragShader, err := utils.LoadShader("shaders/graph.frag")

	if err != nil {
		panic(err)
	}

	vertShader, err := utils.LoadShader("shaders/graph.vert")

	if err != nil {
		panic(err)
	}

	gl.Enable(gl.VERTEX_PROGRAM_POINT_SIZE)

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

func initSineWave() []info.Vec2 {

	points := make([]info.Vec2, 0, width)

	//midPoint := width / 2

	for x := 0; x < width; x++ {
		points = append(points, info.Vec2{
			X: 2*(float32(x)/float32(width)) - 1,
			Y: float32(math.Sin(utils.ToRadians(x))),
		})

	}

	return points
}
