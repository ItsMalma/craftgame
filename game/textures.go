package game

import (
	_ "image/jpeg"
	_ "image/png"
	"minecraft/pkg/gl"
	"minecraft/pkg/glu"

	"bytes"
	"image"
	"image/color"
	"math"
	"os"
)

func LoadTexture(resourceName string, mode int32) (int32, error) {
	var id int32
	gl.GenTextures(1, &id)

	BindTexture(id)

	gl.TexParameteri(gl.Texture2D, gl.TextureMinFilter, mode)
	gl.TexParameteri(gl.Texture2D, gl.TextureMagFilter, mode)

	data, err := os.ReadFile(resourceName)
	if err != nil {
		return 0, err
	}

	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return 0, err
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

	glu.Build2DMipmaps(gl.Texture2D, gl.RGBA, int32(width), int32(height), gl.RGBA, gl.UnsignedByte, &pixels[0])

	return id, nil
}

var lastId int32 = math.MinInt32

func BindTexture(id int32) {
	if id != lastId {
		gl.BindTexture(gl.Texture2D, id)
		lastId = id
	}
}
