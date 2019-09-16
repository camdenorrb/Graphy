package test

import (
	"Graphy/info"
	"Graphy/utils"
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"io/ioutil"
	"log"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// https://kylewbanks.com/blog/tutorial-opengl-with-golang-part-2-drawing-the-game-board

/*
var triangle = []float32 {
	   0.0,  0.5, 0.0, // Top
	  -0.5, -0.5, 0.0, // Left
	   0.5, -0.5, 0.0, // Right
}
*/

var square = []info.Vec3{

	// Half 1
	info.NewVec3(-0.5, 0.5, 0.0),
	info.NewVec3(-0.5, -0.5, 0.0),
	info.NewVec3(0.5, -0.5, 0.0),

	// Half 2
	info.NewVec3(-0.5, 0.5, 0.0),
	info.NewVec3(0.5, 0.5, 0.0),
	info.NewVec3(0.5, -0.5, 0.0),
}

const (
	fps = 60

	rows = 500
	cols = 500

	width  = 1000
	height = 1000
)

type cell struct {
	vao uint32
	pos info.Vec2
}

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func GridMain() {

	window := initGLFW()
	program := initGL()

	cells := initCells()

	r, g, b := utils.NormalizeRGB(153, 211, 205)
	gl.ClearColor(r, g, b, 1)

	for !window.ShouldClose() {

		start := time.Now()
		draw(cells, window, program)
		elapsed := time.Since(start)

		time.Sleep(((1000 / fps) * time.Millisecond) - elapsed)

		// TODO: Calculate how long it takes to draw in order to determine actual FPS
	}

	glfw.Terminate()
}

func initCells() [][]*cell {
	cells := make([][]*cell, rows, rows)
	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			c := newCell(float32(x), float32(y))
			cells[x] = append(cells[x], c)
		}
	}

	return cells
}

func newCell(x, y float32) *cell {

	points := make([]info.Vec3, len(square), len(square))
	copy(points, square)

	for i := 0; i < len(points); i++ {

		point := &points[i]

		sizeX := float32(1.0 / cols)
		sizeY := float32(1.0 / rows)

		posX := x * sizeX
		posY := y * sizeY

		if point.X < 0 {
			point.X = (posX * 2) - 1
		} else {
			point.X = ((posX + sizeX) * 2) - 1
		}

		if point.Y < 0 {
			point.Y = (posY * 2) - 1
		} else {
			point.Y = ((posY + sizeY) * 2) - 1
		}
	}

	return &cell{
		vao: utils.MakeVao(points),
		pos: info.Vec2{X: x, Y: y},
	}
}

func (c *cell) draw() {
	gl.BindVertexArray(c.vao)
	gl.DrawArrays(gl.LINE_STRIP, 0, int32(len(square)))
}

func draw(cells [][]*cell, window *glfw.Window, program uint32) {

	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)

	for _, row := range cells {
		for _, cell := range row {
			cell.draw()
		}
	}

	glfw.PollEvents()
	window.SwapBuffers()
}

func initGL() uint32 {

	if err := gl.Init(); err != nil {
		log.Panicf("Unable to initialize GL %v", err)
	}

	log.Println("OpenGL version", gl.GoStr(gl.GetString(gl.VERSION)))

	fragShader, err := loadShader("shaders/triangle.frag")

	if err != nil {
		panic(err)
	}

	vertShader, err := loadShader("shaders/triangle.vert")

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

func loadShader(filePath string) (uint32, error) {

	ext := strings.TrimPrefix(strings.ToLower(filepath.Ext(filePath)), ".")

	var shaderType uint32

	switch ext {

	case "vert":
		shaderType = gl.VERTEX_SHADER

	case "frag":
		shaderType = gl.FRAGMENT_SHADER

	default:
		log.Panicf("Unknown shader type %s", ext)
	}

	src, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Panicf("Unable to load shader %v", err)
	}

	return utils.CompileShader(string(src)+"\x00", shaderType)
}
