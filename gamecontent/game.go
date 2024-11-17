package gamecontent

import (
	"fmt"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"shoot/assets"
)

const (
	screenWidth = 800
	screenHeight = 800

	meteorSpawnTime = 1 * time.Second

	baseMeteorVelocity = 0.15
	meteorSpeedUpAmount = 0.05
	meteorSpeedUptime = 5 * time.Second

)

type Game struct{
	player       *Player
	meteorSpawnTimer *Timer
	meteor        []*Meteor
	bullets       []*Bullet

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
	g.baseVelocity += meteorSpeedUpAmount
	}

	g.player.Update()

	g.meteorSpawnTimer.Update()
	if g.meteorSpawnTimer.IsReady(){
		g.meteorSpawnTimer.Reset()

		m := NewMeteor(g.baseVelocity)
		g.meteor = append(g.meteor, m)
	}
	for _, m := range g.meteor{
		m.Update()
	}

	for _, b := range g.bullets{
		b.Update()
	}

	for i, m := range g.meteor{
		for j, b := range g.bullets {
			if m.Collider().Intersect(b.Collider()){
				g.meteor = append(g.meteor[:i], g.meteor[i+1:]...)
				g.bullets = append(g.bullets[:j], g.bullets[j+1:]...)
				g.score++
			}
		}
	}

	for _, m := range g.meteor{
		if m.Collider().Intersect(g.player.Collider()){
			g.Reset()
			break
		}
	} 

	return nil
}

func(g *Game) Draw(screen *ebiten.Image){
	g.player.Draw(screen)

	for _, m := range g.meteor{
		m.Draw(screen)
	}

	for _, b := range g.bullets{
		b.Draw(screen)
	}
	text.Draw(screen, fmt.Sprintf("%06d", g.score), assets.ScoreFont, screenWidth/ 2-100, 50, color.White)
}

func(g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int){
	return screenWidth, screenHeight
}
func(g *Game) AddBullet(b *Bullet){
	g.bullets = append(g.bullets, b)
} 

func(g *Game) Reset() {
	g.player = NewPlayer(g)
	g.meteor = nil
	g.bullets = nil
	g.score = 0
	g.meteorSpawnTimer.Reset()
	g.baseVelocity = baseMeteorVelocity
	g.velocityTimer.Reset()
}