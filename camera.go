package main

import "github.com/go-gl/mathgl/mgl32"

type camera struct {
	fov, aspect float32
}

func (c *camera) ray(u, v float32) ray {
	h := tan(c.fov / 2)
	w := h * c.aspect
	return ray{
		direction: mgl32.Vec3{
			w * (2*u - 1),
			h * (2*v - 1),
			-1,
		},
	}
}
