package main

import "github.com/go-gl/mathgl/mgl32"

type material interface {
	scatter(*intersection, mgl32.Vec3) (mgl32.Vec3, mgl32.Vec3)
}

type reflector struct {
	attenuation mgl32.Vec3
}

var _ material = (*reflector)(nil)

func (r *reflector) scatter(i *intersection, in mgl32.Vec3) (mgl32.Vec3, mgl32.Vec3) {
	return in.Sub(i.normal.Mul(2 * in.Dot(i.normal))), r.attenuation
}

type diffuser struct {
	attenuation mgl32.Vec3
}

var _ material = (*diffuser)(nil)

func (d *diffuser) scatter(i *intersection, _ mgl32.Vec3) (mgl32.Vec3, mgl32.Vec3) {
	target := i.position.Add(i.normal).Add(randomInUnitBall())
	return target.Sub(i.position), d.attenuation
}