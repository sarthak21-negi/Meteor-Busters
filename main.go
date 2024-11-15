package main

import( 
	"github.com/hajimehoshi/ebiten/v2"
    "shoot/assets"
)


func main(){

	err := ebiten.RunGame(g)

	if err != nil {
		panic(err)
	}
}
 