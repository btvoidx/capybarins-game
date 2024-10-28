package main

import (
	"errors"
	"fmt"
	"os"

	"game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := &game.Game{}

	err := ebiten.RunGame(game)
	if err != nil && !errors.Is(err, ebiten.Termination) {
		fmt.Fprintf(os.Stderr, "unhandled err: %v\n", err)
		os.Exit(1)
	}
}
