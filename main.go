package main

import (
	"fmt"

	"github.com/go-gl/mathgl/mgl32"
)

func main() {
	nx := 800
	ny := 400
	fmt.Printf("P3\n%d %d\n255\n", nx, ny)
	s := scene{
		top:    blue,
		bottom: white,
		spheres: []sphere{
			{
				center: mgl32.Vec3{0, 0, -5},
				radius: 2,
			},
		},
	}
	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			u := float32(i) / float32(nx)
			v := float32(j) / float32(ny)
			r := ray{
				direction: mgl32.Vec3{
					4*u - 2,
					2*v - 1,
					-1,
				},
			}
			color := s.color(r)
			ir := int(255.99 * color[0])
			ig := int(255.99 * color[1])
			ib := int(255.99 * color[2])
			fmt.Printf("%d %d %d\n", ir, ig, ib)
		}
	}
}
