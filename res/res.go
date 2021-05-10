package res

import (
	"bytes"
	"image"

	"github.com/hajimehoshi/ebiten/v2"

	_ "embed"
)

var (
	//go:embed runner.png
	runner []byte
	//go:embed gnome.png
	gnome []byte
)

func load(pictureData []byte) *ebiten.Image {
	dataReader := bytes.NewReader(pictureData)
	img, _, err := image.Decode(dataReader)
	if err != nil {
		panic(err)
	}
	return ebiten.NewImageFromImage(img)
}

func LoadRunner() *ebiten.Image {
	return load(runner)
}

func LoadGnome() *ebiten.Image {
	return load(gnome)
}
