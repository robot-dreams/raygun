package main

import (
	"github.com/go-gl/mathgl/mgl32"
)

type scene struct {
	top, bottom  mgl32.Vec3
	sceneObjects []sceneObject
}

func (s scene) backgroundColor(r ray) mgl32.Vec3 {
	t := (r.direction.Normalize()[1] + 1) / 2
	return s.top.Mul(t).Add(s.bottom.Mul(1 - t))
}

func (s scene) color(r ray) mgl32.Vec3 {
	for _, o := range s.sceneObjects {
		i := o.intersect(r)
		if i != nil {
			/*
				n := sphere.reflect(r).direction.Normalize()
				return n.Add(white).Mul(0.5)
			*/
			return s.color(i.ray())
		}
	}
	return s.backgroundColor(r)
}

type intersection struct {
	t                float32
	position, normal mgl32.Vec3
}

func (i intersection) ray() ray {
	return ray{
		origin:    i.position,
		direction: i.normal,
	}
}

type sceneObject interface {
	intersect(r ray) *intersection
}
