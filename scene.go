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

func (s scene) color(r ray, remainingBounces int) mgl32.Vec3 {
	if remainingBounces == 0 {
		return s.backgroundColor(r)
	}
	var closest *intersection
	for _, o := range s.sceneObjects {
		i := o.intersect(r)
		if i != nil {
			if closest == nil || i.t < closest.t {
				closest = i
			}
		}
	}
	if closest != nil {
		direction, attenuation := closest.scatter(r.direction)
		return mul3(s.color(direction, remainingBounces-1), attenuation)
	} else {
		return s.backgroundColor(r)
	}
}
