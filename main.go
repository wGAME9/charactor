package main

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

var (
	//go:embed assets/images/runner.png
	runnerImageBytes []byte
)

type game struct {
	player *player
}

func (g *game) Update() error {
	if err := g.player.update(); err != nil {
		return err
	}

	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	g.player.draw(screen)
}

func (g *game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	imageReader := bytes.NewReader(runnerImageBytes)
	runnerImage, _, err := image.Decode(imageReader)
	if err != nil {
		log.Fatalln(err)
	}

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Drawing character")

	runnerStartingX, runnerStartingY := screenWidth/2, screenHeight/2
	runnerStartingX -= frameWidth / 2
	runnerStartingY -= frameHeight / 2

	player := &player{
		imageTiles: ebiten.NewImageFromImage(runnerImage),
		x:          runnerStartingX,
		y:          runnerStartingY,
		count:      0,
		isMoving:   false,
		speed:      speed,
	}

	game := &game{
		player: player,
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
