package unit

import (
	"image"

	"github.com/Frosin/Communis/res"
	"github.com/hajimehoshi/ebiten/v2"
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
	Count,
	backX,
	backY,
	X,
	Y int
}

func (u *Unit) prepereImage() {
	u.unitImage = res.LoadGnome()
}

func NewUnit(backX, backY, x, y int) *Unit {
	u := Unit{
		drawParams: drawParams{
			startFrameOX:     0,
			startFrameOY:     160,
			frameWidth:       38,
			frameHeight:      32,
			frameWidthFloat:  32,
			frameHeightFloat: 32,
			frameNum:         5,
			frameSpeed:       6,
		},
		backX: backX,
		backY: backY,
		X:     x,
		Y:     y,
	}
	u.prepereImage()
	return &u
}

func (u *Unit) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	if u.leftMirror {
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(float64(u.frameWidth), 0)
	}
	op.GeoM.Translate(float64(u.backX+u.X), float64(u.backY+u.Y))

	i := (u.Count / u.frameSpeed) % u.frameNum
	sx, sy := u.startFrameOX+i*u.frameWidth, u.startFrameOY
	screen.DrawImage(u.unitImage.SubImage(image.Rect(sx, sy, sx+u.frameWidth, sy+u.frameHeight)).(*ebiten.Image), op)
}

func (u *Unit) UpdatePosition(count, backX, backY int) {
	u.Count = count
	u.backX = backX
	u.backY = backY

	//u.leftMirror = 0 != consts.LeftMirror&moveKey
}
