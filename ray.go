package main

import "github.com/go-gl/mathgl/mgl32"

type ray struct {
	origin, direction mgl32.Vec3
}

func (r ray) pointAtParameter(t float32) mgl32.Vec3 {
	return r.origin.Add(r.direction.Mul(t))
}

func (r ray) color() mgl32.Vec3 {
	t := (r.direction.Normalize()[1] + 1) / 2
	return white.Mul(1 - t).Add(blue.Mul(t))
}
