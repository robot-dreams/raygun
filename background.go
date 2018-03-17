package main

import "github.com/go-gl/mathgl/mgl32"

type background struct {
	top, bottom mgl32.Vec3
}

func (b background) color(r ray) mgl32.Vec3 {
	t := (r.direction.Normalize()[1] + 1) / 2
	return b.top.Mul(t).Add(b.bottom.Mul(1 - t))
}
