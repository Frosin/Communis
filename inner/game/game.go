package game

import (
	"github.com/Frosin/Communis/inner/consts"
	"github.com/Frosin/Communis/inner/gamemap"
	"github.com/Frosin/Communis/inner/hero"

	"github.com/hajimehoshi/ebiten/v2"
	util "github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	ScreenWidth,
	ScreenHeight int
	Count        int
	Hero         *hero.Hero
	Map          *gamemap.Map
	moveKey      uint8
	heroX, heroY int
	leftIsLast   bool
}

func (g *Game) Update() error {
	g.Count++

	g.checkKeyboard()
	g.Map.Update(g.moveKey, g.heroX, g.heroY, g.Count)
	g.Hero.UpdatePosition(g.Count, g.moveKey)
	g.calculateUnitPosition()

	return nil
}

func (g *Game) checkKeyboard() {
	if util.IsKeyJustPressed(ebiten.KeyDown) {
		g.moveKey |= consts.DownKey
	}
	if util.IsKeyJustReleased(ebiten.KeyDown) {
		g.moveKey &^= consts.DownKey
	}
	if util.IsKeyJustPressed(ebiten.KeyUp) {
		g.moveKey |= consts.UpKey
	}
	if util.IsKeyJustReleased(ebiten.KeyUp) {
		g.moveKey &^= consts.UpKey
	}
	if util.IsKeyJustPressed(ebiten.KeyLeft) {
		g.moveKey |= consts.LeftKey
		g.leftIsLast = true
	}
	if util.IsKeyJustReleased(ebiten.KeyLeft) {
		g.moveKey &^= consts.LeftKey
	}
	if util.IsKeyJustPressed(ebiten.KeyRight) {
		g.moveKey |= consts.RightKey
		g.leftIsLast = false
	}
	if util.IsKeyJustReleased(ebiten.KeyRight) {
		g.moveKey &^= consts.RightKey
	}
	//if finally right pressed
	if 0 != consts.RightKey&g.moveKey {
		g.moveKey &^= consts.LeftMirror
		g.leftIsLast = false
	}
	//if finally lefdt pressed
	if 0 != consts.LeftKey&g.moveKey || g.leftIsLast {
		g.moveKey |= consts.LeftMirror
		g.leftIsLast = true
	}
}

func (g *Game) calculateUnitPosition() {
	backX, backY := g.Map.GetPosition()
	g.heroX = g.ScreenWidth/2 - backX
	g.heroY = g.ScreenHeight/2 - backY
	//debug
	// fmt.Printf("unitX=%d\n", g.heroX)
	// fmt.Printf("unitY=%d\n", g.heroY)
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Map.Draw(screen)
	g.Hero.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.ScreenWidth, g.ScreenHeight
}

func New(screenWidth, screenHeight int) *Game {
	return &Game{
		ScreenWidth:  screenWidth,
		ScreenHeight: screenHeight,
		Hero:         hero.NewHero(screenWidth, screenHeight),
		Map:          gamemap.New(),
	}
}
