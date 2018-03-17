package main

import (
	"fmt"
	"math/rand"

	"github.com/go-gl/mathgl/mgl32"
)

func main() {
	nx := 800
	ny := 400
	subpixelSamples := 100
	fmt.Printf("P3\n%d %d\n255\n", nx, ny)
	s := scene{
		top:    blue,
		bottom: white,
		sceneObjects: []sceneObject{
			sphere{
				center: mgl32.Vec3{0, 0, -1},
				radius: 0.5,
			},
			sphere{
				center: mgl32.Vec3{0, -100.5, -1},
				radius: 100,
			},
		},
	}
	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			color := black
			for k := 0; k < subpixelSamples; k++ {
				u := (float32(i) + rand.Float32()) / float32(nx)
				v := (float32(j) + rand.Float32()) / float32(ny)
				r := ray{
					direction: mgl32.Vec3{
						4*u - 2,
						2*v - 1,
						-1,
					},
				}
				color = color.Add(s.color(r))
			}
			color = color.Mul(1 / float32(subpixelSamples))
			ir := int(255.99 * color[0])
			ig := int(255.99 * color[1])
			ib := int(255.99 * color[2])
			fmt.Printf("%d %d %d\n", ir, ig, ib)
		}
	}
}
