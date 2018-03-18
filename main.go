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
				center: mgl32.Vec3{0, -0.1, -1},
				radius: 0.45,
				m: &diffuser{
					attenuation: mgl32.Vec3{0.9, 0.2, 0.1},
				},
			},
			sphere{
				center: mgl32.Vec3{0, -100.5, -1},
				m: &reflector{
					attenuation: mgl32.Vec3{0.6, 0.7, 0.6},
					blur:        1,
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
				center: mgl32.Vec3{0.5, -0.25, -0.25},
				radius: 0.25,
				m: &refractor{
					attenuation: ones3,
					n:           1.5,
				},
			},
			sphere{
				center:   mgl32.Vec3{0.5, -0.25, -0.25},
				radius:   0.2,
				inverted: true,
				m: &refractor{
					attenuation: ones3,
					n:           1.5,
				},
			},
		},
	}
	c := newCamera(
		mgl32.Vec3{-3, 1, 3},
		mgl32.Vec3{0, 0, -1},
		mgl32.Vec3{0, 1, 0},
		25,
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
			ir := int(255.99 * color[0])
			ig := int(255.99 * color[1])
			ib := int(255.99 * color[2])
			fmt.Printf("%d %d %d\n", ir, ig, ib)
		}
	}
}
