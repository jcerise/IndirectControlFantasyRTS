package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

const (
	DarkGrass1  = 52
	DarGrass2   = 53
	LightGrass1 = 91
	LightGrass2 = 92
	LightGrass3 = 93
	LightGrass4 = 94
)

const (
	TileSize     = 13
	TileSetWidth = 20
)

// GetImageFromTileIndex returns an Ebiten Image from the index on the provided tilesheet (also an Ebiten Image). As
// an example, if zero is provided as tileIndex, the subimage in the provided tilesheet at coordinates (0,0) would be
// returned.
func GetImageFromTileIndex(tileIndex int, tileImage, screen *ebiten.Image) *ebiten.Image {
	sx := (tileIndex % TileSetWidth) * TileSize
	sy := (tileIndex / TileSetWidth) * TileSize

	return tileImage.SubImage(image.Rect(sx, sy, sx+TileSize, sy+TileSize)).(*ebiten.Image)
}
