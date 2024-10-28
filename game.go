package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	w, h int
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen,
		fmt.Sprintf("window size:\n%dx%d", g.w, g.h),
		g.w/2-35, g.h/2-12,
	)
}

func (*Game) Update() error {
	return nil
}

func (g *Game) Layout(w, h int) (int, int) {
	g.w, g.h = w, h
	return w, h
}
