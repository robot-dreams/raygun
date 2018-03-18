package main

import (
	"github.com/go-gl/mathgl/mgl32"
)

type camera struct {
	eye, center, horizontal, vertical mgl32.Vec3

	lensRadius float32
}

func newCamera(eye, center, up mgl32.Vec3, fov, aspect, aperture float32) *camera {
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
		lensRadius: aperture / 2,
	}
}

func (c *camera) ray(u, v float32) ray {
	rd := randomInUnitDisk()
	lensOffset := mgl32.Vec3{
		u * rd[0],
		v * rd[1],
		0,
	}.Mul(c.lensRadius)
	target := c.center
	target = target.Add(c.horizontal.Mul(2*u - 1))
	target = target.Add(c.vertical.Mul(2*v - 1))
	origin := c.eye.Add(lensOffset)
	return ray{
		origin:    origin,
		direction: target.Sub(origin),
	}
}
