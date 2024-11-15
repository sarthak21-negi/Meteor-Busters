package gamecontent

import (
	"fmt"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const (
	screenWidth = 800
	screenHeight = 800

	meteorSpawnTime = 2 * time.Second

	baseMeteorVelocity = 0.15
	meteorSpeedUpAmount = 0.05
	meteorSpeedUptime = 5 * time.Second

)

type Game struct{
	player       *Player
	meteorSpawnTimer *Timer
	meteor        []*Meteor
	bullets       []*bullets

	score int

	baseVelocity  float64
	velocityTimer   *Timer 
}

func NewGame() *Game{
	g := &Game {
		meteorSpawnTimer: NewTimer(meteorSpawnTime),
		baseVelocity: baseMeteorVelocity,
		velocityTimer: NewTimer(meteorSpeedUptime),
	}
	g.player = NewPlayer(g)
	return g
}
func (g *Game) Update() error {
	g.velocityTimer.Update()
	if g.velocityTimer.IsReady() {
	g.velocityTimer.Reset()
	}
	g.player.Update()

}

func(g *Game) Draw(screen *ebiten.Image){
	g.player.Draw(screen)
}

func(g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int){
	return outsideWidth, outsideHeight
}