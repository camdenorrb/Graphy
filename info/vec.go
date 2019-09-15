package info

import "github.com/go-gl/glfw/v3.2/glfw"

type Vec2 struct {
	X float32
	Y float32
}

type Vec3 struct {
	Vec2
	Z float32
}

/*
func NewVec2(x, y float32) Vec2 {
	return Vec2{
		x: x,
		y: y,
	}
}
*/
func NewVec3(x, y, z float32) Vec3 {
	return Vec3{

		Vec2: Vec2{
			X: x,
			Y: y,
		},

		Z: z,
	}
}

func (pos Vec2) NormalizedByWindow(window glfw.Window) *Vec2 {
	width, height := window.GetSize()
	return pos.NormalizedByValues(width, height)
}

func (pos Vec2) NormalizedByValues(windowWidth, windowHeight int) *Vec2 {
	return &Vec2{
		X: 2*pos.X/float32(windowWidth) - 1,
		Y: 2*pos.Y/float32(windowHeight) - 1,
	}
}
