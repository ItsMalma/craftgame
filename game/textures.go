package game

import (
	_ "image/jpeg"
	_ "image/png"
	"log"
	"minecraft/pkg/glu"

	"bytes"
	"image"
	"image/color"
	"math"
	"os"

	"github.com/go-gl/gl/v2.1/gl"
)

func LoadTexture(resourceName string, mode int32) uint32 {
	var id uint32
	gl.GenTextures(1, &id)

	BindTexture(int(id))

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, mode)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, mode)

	data, err := os.ReadFile(resourceName)
	if err != nil {
		log.Fatal(err)
	}

	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}

	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	pixels := make([]int32, width*height)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			color := color.NRGBAModel.Convert(img.At(x, y)).(color.NRGBA)

			alpha := int32(color.A)
			red := int32(color.R)
			green := int32(color.G)
			blue := int32(color.B)

			// Konversi ARGB ke ABGR
			abgr := (alpha << 24) | (blue << 16) | (green << 8) | red

			pixels[(y-bounds.Min.Y)*width+(x-bounds.Min.X)] = abgr
		}
	}

	glu.Build2DMipmaps(gl.TEXTURE_2D, gl.RGBA, int32(width), int32(height), gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(&pixels[0]))

	return id
}

var lastId int = math.MinInt64

func BindTexture(id int) {
	if id != lastId {
		gl.BindTexture(gl.TEXTURE_2D, uint32(id))
		lastId = id
	}
}
