package main

import "github.com/go-gl/mathgl/mgl32"

type ray struct {
	origin, direction mgl32.Vec3
}

func (r ray) pointAtParameter(t float32) mgl32.Vec3 {
	return r.origin.Add(r.direction.Mul(t))
}
