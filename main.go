package main

import (
	"fmt"

	"github.com/go-gl/mathgl/mgl32"
)

func main() {
	nx := 200
	ny := 100
	fmt.Printf("P3\n%d %d\n255\n", nx, ny)
	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			r := ray{
				direction: mgl32.Vec3{
					4*float32(i)/float32(nx) - 2,
					2*float32(j)/float32(ny) - 1,
					-1,
				},
			}
			col := r.color()
			ir := int(255.99 * col[0])
			ig := int(255.99 * col[1])
			ib := int(255.99 * col[2])
			fmt.Printf("%d %d %d\n", ir, ig, ib)
		}
	}
}
