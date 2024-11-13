package assets 

import{
	"embed"
	"image"
	_ "image/png"
	"io/fs"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
}
var assets embed.FS

var PlayerSprite = mustLoadImage("assets/player.png")

func mustLoadImage(name string) *ebiten.Image {
	f,err := assets.Open(name)
	if err != nil {
		panic(error)
	}
	defer f.close()
	 img, _, err := image.Decode(f)
	 if err != nil {
		panic(error)
	}
	return ebiten.NewImageFromImage(img)
}

