package main

import(

	"github.com/hajimehoshi/ebiten/v2"

	"shoot/gamecontent"

)

func main(){
	g := gamecontent.NewGame()
	err := ebiten.RunGame(g)

	if err != nil {
		panic(err)
	}
}
 