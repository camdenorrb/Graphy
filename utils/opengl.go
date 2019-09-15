package utils

import (
	"Graphy/info"
	"fmt"
	"github.com/go-gl/gl/v4.6-core/gl"
	"strings"
)

func NormalizeRGB(ir, ig, ib float32) (or, og, ob float32) {

	or = ir / 255
	og = ig / 255
	ob = ib / 255

	return
}

func MakeVao(points []info.Vec3) (vao uint32) {

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
