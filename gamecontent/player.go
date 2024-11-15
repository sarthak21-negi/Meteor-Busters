package gamecontent

import (
	"math"
	"time"
	"github.com/hajimehoshi/ebiten/v2"
	"shoot/assets"
)

const (
	
)

type Player struct{
	position Vector
	sprite *ebiten.Image
}

func NewPlayer() *Player{
	sprite := assets.PlayerSprite

	bounds := sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	pos := Vector{
		X: screenWidth/2 - halfW,
		Y: screenHeight/2 - halfH,
	}

	return &Player{
		position: pos,
		sprite : sprite,
	}
}

func (p *Player) Update() {

}

func (p *Player) Draw(screen *ebiten.Image){
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.position.X, p.position.Y)
	screen.DrawImage(p.sprite, op)
}