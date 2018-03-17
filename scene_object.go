package main

import "github.com/go-gl/mathgl/mgl32"

type intersection struct {
	t                float32
	position, normal mgl32.Vec3
}

func (i intersection) mirror() ray {
	return ray{
		origin:    i.position,
		direction: i.normal,
	}
}

type sceneObject interface {
	intersect(r ray) *intersection
}
