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

	frameOX     = 0
	frameOY     = 32
	frameWidth  = 32
	frameHeight = 32

	firstLineFrameCount  = 5
	secondLineFrameCount = 8
	thirdLineFrameCount  = 4

	speed = 5
)

var (
	//go:embed assets/images/runner.png
	runnerImageBytes []byte
)

type game struct {
	runnerImage *ebiten.Image
	runnerX     int
	runnerY     int
	count       int
}

func (g *game) Update() error {
	g.count++
	g.count %= 40

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.runnerX -= 2
		if g.runnerX < 0 {
			g.runnerX = 0
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.runnerX += 2
		if g.runnerX > screenWidth-frameWidth {
			g.runnerX = screenWidth - frameWidth
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.runnerY -= 2
		if g.runnerY < 0 {
			g.runnerY = 0
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.runnerY += 2
		if g.runnerY > screenHeight-frameHeight {
			g.runnerY = screenHeight - frameHeight
		}
	}

	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	// Runner image is 256x96, and the character is 32x32.
	//
	// Runner image contain 3 lines of character frames
	// first line has 5 frames, second line has 8 frames, and third line has 4 frames.
	//
	// In this example, we will only use the second line.

	options := &ebiten.DrawImageOptions{}

	options.GeoM.Translate(float64(g.runnerX), float64(g.runnerY))

	// choose which frame to draw
	idx := (g.count / speed) % secondLineFrameCount

	// locate to the frame we want to draw
	// startingX will be frameOX + idx * frameWidth (x coordinate of frame at idx)
	// startingY will be frameOY * 1 (y coordinate of second line)
	startingX, startingY := frameOX+idx*frameWidth, frameOY*1
	endingX, endingY := startingX+frameWidth, startingY+frameHeight

	character := g.runnerImage.SubImage(
		image.Rect(
			startingX, startingY,
			endingX, endingY,
		),
	).(*ebiten.Image)

	screen.DrawImage(character, options)
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

	game := &game{
		runnerImage: ebiten.NewImageFromImage(runnerImage),
		runnerX:     runnerStartingX,
		runnerY:     runnerStartingY,
		count:       0,
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
