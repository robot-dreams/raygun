package main

import "github.com/go-gl/mathgl/mgl32"

type intersection struct {
	t                float32
	position, normal mgl32.Vec3

	m material
}

func (i *intersection) scatter(in mgl32.Vec3) (ray, mgl32.Vec3) {
	direction, attenuation := i.m.scatter(i, in)
	return ray{
		origin:    i.position,
		direction: direction,
	}, attenuation
}

type sceneObject interface {
	intersect(r ray) *intersection
}
