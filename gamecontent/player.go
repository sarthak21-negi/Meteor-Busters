package gamecontent

import (
	"math"
	"time"
	"github.com/hajimehoshi/ebiten/v2"
	"shoot/assets"
)

const (
	shootCoolDown = time.Millisecond *500
	bulletSpawnOffSet = 50.0
)

type Player struct{
	game *Game
	position Vector
	rotation float64
	sprite *ebiten.Image
	shootCoolDown *Timer
}

func NewPlayer(game *Game) *Player{
	sprite := assets.PlayerSprite
	bounds := sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	pos := Vector{
		X: screenWidth/2 - halfW,
		Y: screenHeight/2 - halfH,
	}

	return &Player{
		game: game,
		position: pos,
		rotation: 0,
		sprite : sprite,
		shootCoolDown: NewTimer(shootCoolDown),
	}
}

func (p *Player) Update() {
	speed := math.Pi / float64(ebiten.TPS())

	if ebiten.IsKeyPressed(ebiten.KeyLeft){
		p.rotation -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight){
		p.rotation += speed
	}
	p.shootCoolDown.Update()
	if p.shootCoolDown.IsReady() && ebiten.IsKeyPressed(ebiten.KeySpace){
		p.shootCoolDown.Reset()
	

	bounds := p.sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2
	
	spawnPos := Vector{
		p.position.X + halfW + math.Sin(p.rotation)*bulletSpawnOffSet,
		p.position.Y + halfH + math.Cos(p.rotation)*bulletSpawnOffSet,
	}

	bullets := NewBullet(spawnPos, p.rotation)
	 
	p.game.AddBullet(bullets)
  }
}
func (p *Player) Draw(screen *ebiten.Image){
	bounds := p.sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-halfW, -halfH)
	
	op.GeoM.Rotate(p.rotation)
	op.GeoM.Translate(halfW, halfH)

	op.GeoM.Translate(p.position.X, p.position.Y)

	screen.DrawImage(p.sprite, op)

}

func (p *Player) Collider() Rect{

	bounds := p.sprite.Bounds()

	return NewRect(
		p.position.X,
		p.position.Y,
		float64(bounds.Dx()),
		float64(bounds.Dy()),
	)
}