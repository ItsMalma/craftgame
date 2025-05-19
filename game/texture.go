package game

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/ItsMalma/craftgame/glu"
	"github.com/go-gl/gl/v2.1/gl"
)

func (g *Game) LoadTexture(resourceName string, mode int32) uint32 {
	if texture, ok := g.textures[resourceName]; ok {
		return texture
	}

	var texture uint32
	gl.GenTextures(1, &texture)
	g.textures[resourceName] = texture
	fmt.Printf("Loading texture %s into %d\n", resourceName, texture)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, mode)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, mode)

	resourceFile, err := os.Open(resourceName)
	if err != nil {
		g.onError(err)
		return 0
	}
	defer resourceFile.Close()

	resourceImage, _, err := image.Decode(resourceFile)
	if err != nil {
		g.onError(err)
		return 0
	}

	resourceBounds := resourceImage.Bounds()
	w, h := resourceBounds.Dx(), resourceBounds.Dy()

	pixels := make([]int32, w*h)

	for y := resourceBounds.Min.Y; y < resourceBounds.Max.Y; y++ {
		for x := resourceBounds.Min.X; x < resourceBounds.Max.X; x++ {
			color := color.NRGBAModel.Convert(resourceImage.At(x, y)).(color.NRGBA)
			pixels[(y-resourceBounds.Min.Y)*w+(x-resourceBounds.Min.X)] = (int32(color.A) << 24) | (int32(color.B) << 16) | (int32(color.G) << 8) | int32(color.R)
		}
	}

	glu.Build2DMipmaps(gl.TEXTURE_2D, gl.RGBA, w, h, gl.RGBA, gl.UNSIGNED_BYTE, &pixels[0])

	return texture
}
