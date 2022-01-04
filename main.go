package main

import (
	"MajestyClone/camera"
	"MajestyClone/resources/tilesets"
	"MajestyClone/worldmap"
	"bytes"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"log"
)

const (
	ScreenWidth  int = 320
	ScreenHeight int = 280
	WindowWidth  int = 640
	WindowHeight int = 480
)

var (
	tileImage *ebiten.Image
)

type Game struct {
	worldMap   *worldmap.WorldMap
	gameCamera *camera.GameCamera
}

func init() {
	// Initialize the tileset, by decoding an image from original tilesets byte slice
	img, _, err := image.Decode(bytes.NewReader(tilesets.BasicTiles))
	if err != nil {
		log.Fatal(err)
	}
	tileImage = ebiten.NewImageFromImage(img)
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.worldMap.Render(tileImage, screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func main() {
	worldMap := &worldmap.WorldMap{
		Width:  23,
		Height: 17,
	}
	worldMap.InitializeWorldMap()
	worldMap.MakeGrass()

	gameCamera, _ := camera.NewGameCamera(1, 1, 320, 240)

	g := &Game{
		worldMap:   worldMap,
		gameCamera: gameCamera,
	}

	ebiten.SetWindowSize(WindowWidth, WindowHeight)
	ebiten.SetWindowTitle("Tally Ho!")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
