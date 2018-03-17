package main

import (
	"math"
	"math/rand"

	"github.com/go-gl/mathgl/mgl32"
)

func sqrt(x float32) float32 {
	return float32(math.Sqrt(float64(x)))
}

func randomInUnitBall() mgl32.Vec3 {
	for {
		candidate := mgl32.Vec3{
			rand.Float32(),
			rand.Float32(),
			rand.Float32(),
		}
		candidate = candidate.Mul(2)
		candidate = candidate.Sub(ones3)
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
