package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"

	"github.com/iliaszh/raytracer/geometry"
)

const (
	imgWidth  = 1280
	imgHeight = 720
)

func main() {
	img := image.NewRGBA(image.Rect(0, 0, imgWidth, imgHeight))

	for y := 0; y < imgHeight; y++ {
		for x := 0; x < imgWidth; x++ {

			img.Set(x, y, color.RGBA{
				R: 255,
				G: uint8((y << 8) / imgHeight),
				B: uint8((x << 8) / imgWidth),
				A: 255,
			})
		}
	}

	file, errCreate := os.Create("output.png")
	if errCreate != nil {
		log.Fatalf("creating file: %v", errCreate)
	}

	errEncode := png.Encode(file, img)
	if errEncode != nil {
		log.Fatalf("encoding image to png: %v", errEncode)
	}
}

func castRay[F geometry.Float](ray geometry.Ray[F], sphere geometry.Sphere[F]) color.Color {
	var (
		backgroundColor = color.RGBA{
			R: 100,
			G: 100,
			B: 100,
			A: 255,
		}
		sphereColor = color.RGBA{
			R: 200,
			G: 180,
			B: 150,
			A: 255,
		}
	)

	_, intersects := sphere.Intersect(ray)
	if intersects {
		return sphereColor
	}

	return backgroundColor
}
