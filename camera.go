package main

import (
	"github.com/go-gl/mathgl/mgl32"
)

type camera struct {
	eye, center, horizontal, vertical mgl32.Vec3
}

func newCamera(eye, center, up mgl32.Vec3, fov, aspect float32) *camera {
	look := center.Sub(eye)
	h := tan(fov/2) * look.Len()
	w := h * aspect
	xAxis := look.Cross(up).Normalize()
	yAxis := xAxis.Cross(look).Normalize()
	return &camera{
		eye:        eye,
		center:     center,
		horizontal: xAxis.Mul(w),
		vertical:   yAxis.Mul(h),
	}
}

func (c *camera) ray(u, v float32) ray {
	target := c.center
	target = target.Add(c.horizontal.Mul(2*u - 1))
	target = target.Add(c.vertical.Mul(2*v - 1))
	return ray{
		origin:    c.eye,
		direction: target.Sub(c.eye),
	}
}
