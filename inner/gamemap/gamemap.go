package gamemap

import (
	"fmt"
	"image/color"

	"github.com/Frosin/Communis/inner/consts"
	"github.com/Frosin/Communis/inner/gamemap/unit"
	"github.com/Frosin/Communis/inner/limits"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	backWidth = 400
)

type Map struct {
	backX,
	backY,
	backWidth int
	limits *limits.Limits
	unit   *unit.Unit
}

func New(limits *limits.Limits) *Map {
	unit := unit.NewUnit(0, 0, 20, 100, limits)

	newMap := Map{
		backX:     0,
		backY:     0,
		backWidth: backWidth,
		limits:    limits,
		unit:      unit,
	}
	limits.SetLimit(50, 70, 30, 50)

	limits.SetLimit(100, 70, 10, 50)

	limits.SetLimit(120, 70, 10, 50)

	limits.SetLimit(100, 150, 30, 50)

	return &newMap
}

func (m *Map) Update(moveKey uint8, heroX, heroY, count int) {
	m.unit.UpdatePosition(count, m.backX, m.backY)

	if 0 != consts.UpKey&moveKey && m.limits.IsValidPosition(heroX, heroY-1) {
		m.backY++
	}
	if 0 != consts.DownKey&moveKey && m.limits.IsValidPosition(heroX, heroY+1) {
		m.backY--
	}
	if 0 != consts.RightKey&moveKey && m.limits.IsValidPosition(heroX+1, heroY) {
		m.backX--
	}
	if 0 != consts.LeftKey&moveKey && m.limits.IsValidPosition(heroX-1, heroY) {
		m.backX++
	}
	//debug
	curX, curY := ebiten.CursorPosition()
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) &&
		m.limits.IsValidPosition(curX-m.backX, curY-m.backY) {
		m.unit.Logic.SetTarget(curX-m.backX, curY-m.backY)
		fmt.Println("set target:", curX-m.backX, curY-m.backY)
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
		float64(m.backY+70),
		float64(10),
		float64(50),
		color.RGBA{200, 0, 0, 255},
	)

	ebitenutil.DrawRect(
		screen,
		float64(m.backX+120),
		float64(m.backY+70),
		float64(10),
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

	m.unit.Draw(screen)
}

func (m *Map) GetPosition() (int, int) {
	return m.backX, m.backY
}
