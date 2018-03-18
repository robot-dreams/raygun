package main

import (
	"math"
	"math/rand"

	"github.com/go-gl/mathgl/mgl32"
)

func pow(a, b float32) float32 {
	return float32(math.Pow(float64(a), float64(b)))
}

func sqrt(x float32) float32 {
	return float32(math.Sqrt(float64(x)))
}

func tan(x float32) float32 {
	return float32(math.Tan(math.Pi * float64(x) / 180))
}

func randomInUnitLine() float32 {
	return 2*rand.Float32() - 1
}

func randomInUnitDisk() mgl32.Vec2 {
	for {
		candidate := mgl32.Vec2{
			randomInUnitLine(),
			randomInUnitLine(),
		}
		if candidate.Len() <= 1 {
			return candidate
		}
	}
}

func randomInUnitBall() mgl32.Vec3 {
	for {
		candidate := mgl32.Vec3{
			randomInUnitLine(),
			randomInUnitLine(),
			randomInUnitLine(),
		}
		if candidate.Len() <= 1 {
			return candidate
		}
	}
}

func gamma2(color mgl32.Vec3) mgl32.Vec3 {
	return mgl32.Vec3{
		sqrt(color[0]),
		sqrt(color[1]),
		sqrt(color[2]),
	}
}

func mul3(u, v mgl32.Vec3) mgl32.Vec3 {
	return mgl32.Vec3{
		u[0] * v[0],
		u[1] * v[1],
		u[2] * v[2],
	}
}

func reflect(v, normal mgl32.Vec3) mgl32.Vec3 {
	return v.Sub(normal.Mul(2 * v.Dot(normal)))
}

// Parameters:
//     v: direction of incoming ray
//     n: outward normal of surface
//     eta: ratio of refractive indexes, outside / inside
//
// Preconditions:
//     n is normalized
func refract(v, n mgl32.Vec3, eta float32) mgl32.Vec3 {
	vHat := v.Normalize()
	c1 := vHat.Dot(n)
	// TODO: Do we need to explicitly handle the case where the dot product is
	// 0 or very close?
	if c1 < 0 {
		// Incoming direction is from outside the surface.
		c1 = -c1
	} else {
		// Incoming direction is from inside the surface.
		n = n.Mul(-1)
		eta = 1 / eta
	}
	discriminant := 1 - eta*eta*(1-c1*c1)
	if discriminant < 0 {
		return reflect(v, n)
	} else {
		return vHat.Mul(eta).Add(n.Mul(eta*c1 - sqrt(discriminant)))
	}
}
