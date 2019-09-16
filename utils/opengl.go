package utils

import (
	"Graphy/info"
	"fmt"
	"github.com/go-gl/gl/v4.6-core/gl"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

func NormalizeRGB(ir, ig, ib float32) (or, og, ob float32) {

	or = ir / 255
	og = ig / 255
	ob = ib / 255

	return
}

func MakeVaoByVec2(points []info.Vec2) (vao uint32) {

	var vbo uint32

	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*(len(points)*2), gl.Ptr(points), gl.STATIC_DRAW)

	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 2, gl.FLOAT, false, 0, nil)

	return
}
func MakeVaoByVec3(points []info.Vec3) (vao uint32) {

	var vbo uint32

	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*(len(points)*3), gl.Ptr(points), gl.STATIC_DRAW)

	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	return
}

func LoadShader(filePath string) (uint32, error) {

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

	return CompileShader(string(src)+"\x00", shaderType)
}

func CompileShader(source string, shaderType uint32) (shader uint32, err error) {

	var status int32

	shader = gl.CreateShader(shaderType)

	glSrc, freeFn := gl.Strs(source)

	gl.ShaderSource(shader, 1, glSrc, nil)
	gl.CompileShader(shader)

	freeFn()

	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)

	if status == gl.FALSE {

		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		logs := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(logs))

		err = fmt.Errorf("failed to compile %v: %v", source, logs)
	}

	return
}
