package entity

import (
	"cheezewiz/config"
	"cheezewiz/internal/component"
	"cheezewiz/pkg/ecs"
	"cheezewiz/pkg/gamemath"
	"fmt"
)

// Level is a representataion of level data for efficient lookup
type Level struct {
	tiles [][]*component.Tile
}

func (l Level) ToString() string {
	var result string
	for y, row := range l.tiles {
		for x, tile := range row {
			// result = append(result, len(tile.Actors))
			result = fmt.Sprintf("%d, %d, - actors: %d", x, y, len(tile.Actors))
		}
	}
	return result
}

func (l Level) GetTile(column int, row int) *component.Tile {
	return l.tiles[row][column]
}

func (l Level) MoveActor(origin component.Position, destination component.Position, handle ecs.EntityHandle, entity ecs.Entity) {
	// o := l.tiles[int(destination.Y)][int(destination.X)].GetTile()
	// d := l.tiles[int(destination.Y)][int(destination.X)].GetTile()

	originTileCoord := gamemath.Vector([]float64{origin.X, origin.Y}).ToTileCoord(config.Get().TileSize)
	originTile := l.GetTile(int(originTileCoord[0]), int(originTileCoord[1]))

	destinationTileCoord := gamemath.Vector([]float64{destination.X, destination.Y}).ToTileCoord(config.Get().TileSize)
	destinationTile := l.GetTile(int(destinationTileCoord[0]), int(destinationTileCoord[1]))

	originTile.Exit(handle)
	destinationTile.Enter(handle, entity)
}
