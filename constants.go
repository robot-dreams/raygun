package main

import "github.com/go-gl/mathgl/mgl32"

const (
	maxBounces = 50
)

var (
	zero3 = mgl32.Vec3{0, 0, 0}
	ones3 = mgl32.Vec3{1, 1, 1}

	black = zero3
	white = ones3
	blue  = mgl32.Vec3{0.5, 0.7, 1}
)
