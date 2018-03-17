package main

import (
	"github.com/go-gl/mathgl/mgl32"
)

type scene struct {
	top, bottom mgl32.Vec3
	spheres     []sphere
}

func (s scene) backgroundColor(r ray) mgl32.Vec3 {
	t := (r.direction.Normalize()[1] + 1) / 2
	return s.top.Mul(t).Add(s.bottom.Mul(1 - t))
}

func (s scene) color(r ray) mgl32.Vec3 {
	for _, sphere := range s.spheres {
		if sphere.intersects(r) {
			/*
				n := sphere.reflect(r).direction.Normalize()
				return n.Add(white).Mul(0.5)
			*/
			return s.color(sphere.reflect(r))
		}
	}
	return s.backgroundColor(r)
}
