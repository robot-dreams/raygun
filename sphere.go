package main

import (
	"github.com/go-gl/mathgl/mgl32"
)

type sphere struct {
	center mgl32.Vec3
	radius float32

	inverted bool
	m        material
}

var _ sceneObject = sphere{}

func (s sphere) intersect(r ray) *intersection {
	offset := r.origin.Sub(s.center)
	a := r.direction.Dot(r.direction)
	b := r.direction.Dot(offset)
	c := offset.Dot(offset) - s.radius*s.radius
	discriminant := b*b - a*c
	if discriminant <= 0 {
		return nil
	}
	for _, t := range []float32{
		(-b - sqrt(discriminant)) / a,
		(-b + sqrt(discriminant)) / a,
	} {
		// TODO: Come up with a better way to deal with the case where the
		// "surface tangent" is slightly inside the sphere due to numerical
		// accuracy issues.
		if t*r.direction.Len() > 1e-3 {
			position := r.pointAtParameter(t)
			normal := position.Sub(s.center).Normalize()
			if s.inverted {
				normal = normal.Mul(-1)
			}
			return &intersection{
				t:        t,
				position: position,
				normal:   normal,

				m: s.m,
			}
		}
	}
	return nil
}
