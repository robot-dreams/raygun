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

func (i intersection) diffuse() ray {
	target := i.position.Add(i.normal).Add(randomInUnitBall())
	return ray{
		origin:    i.position,
		direction: target.Sub(i.position),
	}
}

type sceneObject interface {
	intersect(r ray) *intersection
}
