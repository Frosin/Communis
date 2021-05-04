package gamemap

import (
	"image/color"

	"github.com/Frosin/Communis/inner/consts"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	backWidth = 200
)

type limitParams struct {
	x, y, w, h int
}

type Map struct {
	backX,
	backY,
	backWidth int
	limits []limitParams
}

func New() *Map {
	newMap := Map{
		0, 0, backWidth,
		make([]limitParams, 0),
	}
	newMap.SetLimit(50, 70, 30, 50)
	newMap.SetLimit(100, 150, 30, 50)
	return &newMap
}

func (m *Map) Update(moveKey uint8, heroX, heroY int) {
	if 0 != consts.UpKey&moveKey && m.isHeroValidPosition(heroX, heroY-1) {
		m.backY++
	}
	if 0 != consts.DownKey&moveKey && m.isHeroValidPosition(heroX, heroY+1) {
		m.backY--
	}
	if 0 != consts.RightKey&moveKey && m.isHeroValidPosition(heroX+1, heroY) {
		m.backX--
	}
	if 0 != consts.LeftKey&moveKey && m.isHeroValidPosition(heroX-1, heroY) {
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

	ebitenutil.DrawRect(
		screen,
		float64(m.backX+50),
		float64(m.backY+70),
		float64(30),
		float64(50),
		color.RGBA{200, 0, 0, 255},
	)

	ebitenutil.DrawRect(
		screen,
		float64(m.backX+100),
		float64(m.backY+150),
		float64(30),
		float64(50),
		color.RGBA{0, 0, 200, 200},
	)

}

func (m *Map) GetPosition() (int, int) {
	return m.backX, m.backY
}

func (m *Map) isHeroValidPosition(uX, uY int) bool {
	for _, l := range m.limits {
		if uX > l.x &&
			uX < l.x+l.w &&
			uY > l.y &&
			uY < l.y+l.h {
			return false
		}
	}
	return true
}

func (m *Map) SetLimit(rectX, rectY, rectWidth, rectHeight int) {
	m.limits = append(m.limits, limitParams{rectX, rectY, rectWidth, rectHeight})
}
