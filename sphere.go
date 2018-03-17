package main

import (
	"github.com/go-gl/mathgl/mgl32"
)

type sphere struct {
	center mgl32.Vec3
	radius float32
}

func (s sphere) intersections(r ray) []mgl32.Vec3 {
	offset := r.origin.Sub(s.center)
	a := r.direction.Dot(r.direction)
	b := 2 * r.direction.Dot(offset)
	c := offset.Dot(offset) - s.radius*s.radius
	discriminant := b*b - 4*a*c
	if discriminant <= 0 {
		return nil
	}
	var result []mgl32.Vec3
	for _, t := range []float32{
		(-b - sqrt(discriminant)) / (2 * a),
		(-b + sqrt(discriminant)) / (2 * a),
	} {
		// TODO: Come up with a better way to deal with the case where the
		// "surface tangent" is slightly inside the sphere due to numerical
		// accuracy issues.
		if t*r.direction.Len() > 1e-3 {
			result = append(result, r.pointAtParameter(t))
		}
	}
	return result
}

func (s sphere) intersects(r ray) bool {
	return len(s.intersections(r)) > 0
}

// Precondition: s.intersects(r)
func (s sphere) reflect(r ray) ray {
	tangent := s.intersections(r)[0]
	return ray{
		origin:    tangent,
		direction: tangent.Sub(s.center),
	}
}
