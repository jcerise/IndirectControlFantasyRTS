package worldmap

import (
	"MajestyClone/ui"
	"github.com/hajimehoshi/ebiten/v2"
	_ "image/png"
	"math/rand"
)

type CoordinatePair struct {
	X int
	Y int
}

type Tile struct {
	SpriteIndex int
	Visible     bool
	Visited     bool
	Explored    bool
	X           int
	Y           int
}

type WorldMap struct {
	Width  int
	Height int
	Tiles  [][]*Tile
}

func (wm *WorldMap) InitializeWorldMap() {
	wm.Tiles = make([][]*Tile, wm.Width+1)
	for i := range wm.Tiles {
		wm.Tiles[i] = make([]*Tile, wm.Height+1)
	}
}

// MakeGrass sets every tile on the map to grass. This is a test method.
func (wm *WorldMap) MakeGrass() {
	for x := 0; x <= wm.Width; x++ {
		for y := 0; y <= wm.Height; y++ {
			var spriteIndex int
			switch rand.Intn(6) {
			case 0:
				spriteIndex = ui.DarkGrass1
			case 1:
				spriteIndex = ui.DarGrass2
			case 2:
				spriteIndex = ui.LightGrass1
			case 3:
				spriteIndex = ui.LightGrass2
			case 4:
				spriteIndex = ui.LightGrass3
			case 5:
				spriteIndex = ui.LightGrass4
			}

			tile := Tile{
				SpriteIndex: spriteIndex,
				Visible:     true,
				Visited:     true,
				Explored:    true,
				X:           x,
				Y:           y,
			}
			wm.Tiles[x][y] = &tile
		}
	}
}

func (wm *WorldMap) Render(tileImage *ebiten.Image, screen *ebiten.Image) {
	for x := 0; x <= wm.Width; x++ {
		for y := 0; y <= wm.Height; y++ {
			tile := wm.Tiles[x][y]

			if tile.Visible {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(x*ui.TileSize), float64(y*ui.TileSize))
				tileImage := ui.GetImageFromTileIndex(tile.SpriteIndex, tileImage, screen)
				screen.DrawImage(tileImage, op)
			}
		}
	}
}
