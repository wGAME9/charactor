package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	frameOX     = 0
	frameOY     = 32
	frameWidth  = 32
	frameHeight = 32

	firstLineFrameCount  = 5
	secondLineFrameCount = 8
	thirdLineFrameCount  = 4

	speed = 5
)

type player struct {
	x, y       int
	imageTiles *ebiten.Image

	count     int
	isMoving  bool
	direction int
	speed     int
}

func (p *player) update() error {
	p.isMoving = false
	p.count++
	p.count %= 40

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.isMoving = true
		p.direction = 1
		p.x -= 2
		if p.x < 0 {
			p.x = 0
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.isMoving = true
		p.direction = 0
		p.x += 2
		if p.x > screenWidth-frameWidth {
			p.x = screenWidth - frameWidth
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		p.isMoving = true
		p.y -= 2
		if p.y < 0 {
			p.y = 0
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		p.isMoving = true
		p.y += 2
		if p.y > screenHeight-frameHeight {
			p.y = screenHeight - frameHeight
		}
	}

	return nil
}

func (p *player) draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}

	if p.direction == 1 {
		options.GeoM.Scale(-1, 1)
		options.GeoM.Translate(frameWidth, 0)
	}

	options.GeoM.Translate(float64(p.x), float64(p.y))

	screen.DrawImage(p.getPlayerImage(), options)
}

func (p *player) getPlayerImage() *ebiten.Image {
	switch {
	case p.isMoving:
		return p.getPlayerMovingImage()
	default:
		return p.getPlayerIdleImage()
	}
}

func (p *player) getPlayerIdleImage() *ebiten.Image {
	idx := (p.count / 10) % firstLineFrameCount

	startingX, startingY := frameOX+idx*frameWidth, 0
	endingX, endingY := startingX+frameWidth, startingY+frameHeight

	character := p.imageTiles.SubImage(
		image.Rect(
			startingX, startingY,
			endingX, endingY,
		),
	).(*ebiten.Image)

	return character
}

func (p *player) getPlayerMovingImage() *ebiten.Image {
	idx := (p.count / speed) % secondLineFrameCount

	startingX, startingY := frameOX+idx*frameWidth, frameOY*1
	endingX, endingY := startingX+frameWidth, startingY+frameHeight

	character := p.imageTiles.SubImage(
		image.Rect(
			startingX, startingY,
			endingX, endingY,
		),
	).(*ebiten.Image)

	return character
}
