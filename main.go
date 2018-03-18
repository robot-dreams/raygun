package main

import (
	"fmt"
	"math/rand"

	"github.com/go-gl/mathgl/mgl32"
)

func main() {
	nx := 200
	ny := 100
	subpixelSamples := 100
	fmt.Printf("P3\n%d %d\n255\n", nx, ny)
	s := scene{
		top:    blue,
		bottom: white,
		sceneObjects: []sceneObject{
			sphere{
				center: mgl32.Vec3{0, 0, -1},
				radius: 0.5,
				m: &diffuser{
					attenuation: mgl32.Vec3{0.8, 0.3, 0.3},
				},
			},
			sphere{
				center: mgl32.Vec3{0, -100.5, -1},
				m: &diffuser{
					attenuation: mgl32.Vec3{0.8, 0.8, 0},
				},
				radius: 100,
			},
			sphere{
				center: mgl32.Vec3{1, 0, -1},
				radius: 0.5,
				m: &reflector{
					attenuation: mgl32.Vec3{0.2, 0.6, 0.8},
					blur:        0.2,
				},
			},
			sphere{
				center: mgl32.Vec3{-1, 0, -1},
				radius: 0.5,
				m: &refractor{
					attenuation: ones3,
					n:           1.5,
				},
			},
			sphere{
				center:   mgl32.Vec3{-1, 0, -1},
				radius:   0.4,
				inverted: true,
				m: &refractor{
					attenuation: ones3,
					n:           1.5,
				},
			},
		},
	}
	c := newCamera(
		mgl32.Vec3{3, 3, 2},
		mgl32.Vec3{0, 0, -1},
		mgl32.Vec3{0, 1, 0},
		20,
		float32(nx)/float32(ny),
		0)
	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			color := black
			for k := 0; k < subpixelSamples; k++ {
				u := (float32(i) + rand.Float32()) / float32(nx)
				v := (float32(j) + rand.Float32()) / float32(ny)
				r := c.ray(u, v)
				color = color.Add(s.color(r, maxBounces))
			}
			color = color.Mul(1 / float32(subpixelSamples))
			// color = gamma2(color)
			ir := int(255.99 * color[0])
			ig := int(255.99 * color[1])
			ib := int(255.99 * color[2])
			fmt.Printf("%d %d %d\n", ir, ig, ib)
		}
	}
}
