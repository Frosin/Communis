package gamemap

import (
	"image/color"

	"github.com/Frosin/Communis/inner/consts"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const backWidth = 200

type Map struct {
	backX,
	backY,
	backWidth int
}

func New() *Map {
	return &Map{
		0, 0, backWidth,
	}
}

func (m *Map) Update(moveKey uint8) {
	if 0 != consts.UpKey&moveKey {
		m.backY++
	}
	if 0 != consts.DownKey&moveKey {
		m.backY--
	}
	if 0 != consts.RightKey&moveKey {
		m.backX--
	}
	if 0 != consts.LeftKey&moveKey {
		m.backX++
	}
}

func (m *Map) Draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(
		screen,
		float64(m.backX),
		float64(m.backY),
		float64(m.backWidth),
		float64(m.backWidth),
		color.RGBA{100, 155, 0, 255},
	)
}

func (m *Map) GetPosition() (int, int) {
	return m.backX, m.backY
}
