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
)

type Game struct{
	player       *Player
}
