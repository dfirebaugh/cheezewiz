package system

import (
	"cheezewiz/config"
	"cheezewiz/internal/component"
	"cheezewiz/internal/entity"
	"cheezewiz/internal/tag"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
)

type WorldViewPortLocation struct {
	playerQuery *query.Query
	query       *query.Query
}

func NewWorldViewPortLocation() *WorldViewPortLocation {
	return &WorldViewPortLocation{
		playerQuery: query.NewQuery(filter.Contains(
			tag.Player,
		)),
		query: query.NewQuery(filter.Contains(
			entity.WorldViewPortTag,
		)),
	}
}

func (worldViewPortLocation *WorldViewPortLocation) Update(w donburi.World) {
	initialPlayer, _ := worldViewPortLocation.playerQuery.FirstEntity(w)

	playerPosition := component.GetPosition(initialPlayer)

	worldViewPort, _ := worldViewPortLocation.query.FirstEntity(w)

	worldViewPortPos := component.GetPosition(worldViewPort)

	worldViewPortPos.X = getWorldViewCenterLocation(playerPosition.X, config.Get().Window.Height) + 20
	worldViewPortPos.Y = 0
}

func getWorldViewCenterLocation(coordinate float64, windowDim int) float64 {
	return coordinate - float64(windowDim/config.Get().ScaleFactor)/2
}
