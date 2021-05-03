package hero

import (
	"image"
	"image/color"
	_ "image/png"
	"log"
	"os"
	"test/inner/consts"

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

func (u *Hero) prepereImage() {
	imgFile, err := os.Open("./inner/hero/runner.png")
	if err != nil {
		panic(err)
	}

	img, _, err := image.Decode(imgFile)
	if err != nil {
		log.Fatal(err)
	}
	u.unitImage = ebiten.NewImageFromImage(img)
}

func NewHero(screenWidth, screenHeight int) *Hero {
	u := Hero{
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
		//Count:        count,
		ScreenWidth:  screenWidth,
		ScreenHeight: screenHeight,
	}
	u.prepereImage()
	return &u
}

func (u *Hero) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	if u.leftMirror {
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(float64(u.frameWidth), 0)
	}
	finalX := float64(u.ScreenWidth)/2 - u.frameWidthFloat/2
	finalY := float64(u.ScreenHeight)/2 - u.frameHeightFloat
	op.GeoM.Translate(finalX, finalY)

	i := (u.Count / u.frameSpeed) % u.frameNum
	sx, sy := u.startFrameOX+i*u.frameWidth, u.startFrameOY
	screen.DrawImage(u.unitImage.SubImage(image.Rect(sx, sy, sx+u.frameWidth, sy+u.frameHeight)).(*ebiten.Image), op)
	//debug point
	ebitenutil.DrawRect(
		screen,
		finalX+u.frameWidthFloat/2,
		finalY+u.frameHeightFloat,
		2,
		2,
		color.RGBA{255, 0, 0, 255},
	)
}

func (u *Hero) UpdatePosition(count int, moveKey uint8) {
	u.Count = count
	u.leftMirror = 0 != consts.LeftMirror&moveKey
}
