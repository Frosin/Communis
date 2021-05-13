package unit

import (
	"image"
	"image/color"

	"github.com/Frosin/Communis/inner/limits"
	"github.com/Frosin/Communis/res"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type drawParams struct {
	frameWidthFloat  float64
	frameHeightFloat float64
	startFrameOX     int
	startFrameOY     int
	frameWidth       int
	frameHeight      int
	frameNum         int
	frameSpeed       int
	unitImage        *ebiten.Image
}

type unitPosition struct {
	leftMirror bool
}

type Unit struct {
	drawParams
	unitPosition
	unitLogic
	Count,
	backX,
	backY,
	X,
	Y int
}

func (u *Unit) prepereImage() {
	u.unitImage = res.LoadGnome()
}

func NewUnit(backX, backY, x, y int, limits *limits.Limits) *Unit {
	u := Unit{
		drawParams: drawParams{
			startFrameOX:     0,
			startFrameOY:     165,
			frameWidth:       38,
			frameHeight:      24,
			frameWidthFloat:  38,
			frameHeightFloat: 24,
			frameNum:         1, //5,
			frameSpeed:       6,
		},
		backX:     backX,
		backY:     backY,
		X:         x,
		Y:         y,
		unitLogic: newLogic(limits),
	}
	u.SetTarget(200, 80)
	u.prepereImage()
	return &u
}

func (u *Unit) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	if u.leftMirror {
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(float64(u.frameWidth), 0)
	}
	op.GeoM.Translate(float64(u.backX+u.X)-u.frameWidthFloat/2, float64(u.backY+u.Y)-u.frameHeightFloat)

	i := (u.Count / u.frameSpeed) % u.frameNum
	sx, sy := u.startFrameOX+i*u.frameWidth, u.startFrameOY
	screen.DrawImage(u.unitImage.SubImage(image.Rect(sx, sy, sx+u.frameWidth, sy+u.frameHeight)).(*ebiten.Image), op)
	//debug point
	ebitenutil.DrawRect(
		screen,
		float64(u.backX+u.X),
		float64(u.backY+u.Y),
		2,
		2,
		color.RGBA{50, 0, 255, 255},
	)
}

func (u *Unit) UpdatePosition(count, backX, backY int) {
	u.Count = count
	u.backX = backX
	u.backY = backY

	if u.haveTargets() {
		u.X, u.Y = u.NextXY(u.X, u.Y)
	}

	//u.calculateLogic()

	//u.leftMirror = 0 != consts.LeftMirror&moveKey
}
