package gamecontent

import (
	"math"
	"math/rand"
	"shoot/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	minRotationSpeed = -0.02
	maxRotationSpeed = 0.02
)

type Meteor struct {
	position Vector
	rotation float64
	rotationSpeed float64
	movement Vector
	sprite *ebiten.Image
}

func NewMeteor(baseVelocity float64) *Meteor{
	target := Vector{
		X: screenWidth / 2,
		Y: screenHeight /2,
	}

	r := screenWidth / 2

	angle := rand.Float64() * 2 * math.Pi

	pos := Vector{
		X: target.X + math.Cos(angle)*float64(r),
		Y: target.Y + math.Sin(angle)*float64(r),
	}

	velocity := baseVelocity + rand.Float64()*1.5

	direction := Vector{
		X: target.X - pos.X,
		Y: target.Y - pos.Y,
	}

	normalizedDirection := direction.Normalize()

	movement := Vector{
		X: normalizedDirection.X * velocity,
		Y: normalizedDirection.Y * velocity,
	}

	sprite := assets.MeteorSprite[rand.Intn(len(assets.MeteorSprite))]

	m := &Meteor{
		position: pos,
		movement: movement,
		rotationSpeed: minRotationSpeed + rand.Float64()*(maxRotationSpeed - minRotationSpeed),
		sprite: sprite,
	}

	return m

}

func (m *Meteor) Update(){
	m.position.X += m.movement.X 
	m.position.Y += m.movement.Y
	m.rotation += m.rotationSpeed 
}

func (m *Meteor) Draw(screen *ebiten.Image){

	bounds := m.sprite.Bounds()

	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(-halfW, -halfH)

	op.GeoM.Rotate(m.rotation)
	op.GeoM.Translate(halfW,halfH)
	op.GeoM.Translate(m.position.X,m.position.Y)
	screen.DrawImage(m.sprite, op)

}
func (m *Meteor) Collider() Rect{

	bounds := m.sprite.Bounds()
	
	return NewRect(
		m.position.X,
		m.position.Y,
		float64(bounds.Dx()),
		float64(bounds.Dy()),
	)
}