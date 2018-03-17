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
