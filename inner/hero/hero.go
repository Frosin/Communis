package hero

import (
	"image"
	"image/color"
	_ "image/png"

	"github.com/Frosin/Communis/inner/consts"
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

type Hero struct {
	drawParams
	unitPosition
	Count int
	ScreenWidth,
	ScreenHeight int
}

func (h *Hero) prepereImage() {
	h.unitImage = res.LoadRunner()
}

func NewHero(screenWidth, screenHeight int) *Hero {
	h := Hero{
		drawParams: drawParams{
			startFrameOX:     0,
			startFrameOY:     32,
			frameWidth:       32,
			frameHeight:      32,
			frameWidthFloat:  32,
			frameHeightFloat: 32,
			frameNum:         8,
			frameSpeed:       6,
		},
		ScreenWidth:  screenWidth,
		ScreenHeight: screenHeight,
	}
	h.prepereImage()
	return &h
}

func (h *Hero) Draw(screen *ebiten.Image) {
	//debug
	//println("count=", h.Count)
	op := &ebiten.DrawImageOptions{}
	if h.leftMirror {
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(float64(h.frameWidth), 0)
	}
	finalX := float64(h.ScreenWidth)/2 - h.frameWidthFloat/2
	finalY := float64(h.ScreenHeight)/2 - h.frameHeightFloat
	op.GeoM.Translate(finalX, finalY)

	i := (h.Count / h.frameSpeed) % h.frameNum
	sx, sy := h.startFrameOX+i*h.frameWidth, h.startFrameOY
	screen.DrawImage(h.unitImage.SubImage(image.Rect(sx, sy, sx+h.frameWidth, sy+h.frameHeight)).(*ebiten.Image), op)
	//debug point
	ebitenutil.DrawRect(
		screen,
		finalX+h.frameWidthFloat/2,
		finalY+h.frameHeightFloat,
		2,
		2,
		color.RGBA{255, 0, 0, 255},
	)
}

func (h *Hero) UpdatePosition(count int, moveKey uint8) {
	if moveKey != 0 && moveKey != 16 {
		h.Count = count
		h.startFrameOY = 32
		h.leftMirror = 0 != consts.LeftMirror&moveKey
	} else {
		h.Count = 0
		h.startFrameOY = 0
	}
}
