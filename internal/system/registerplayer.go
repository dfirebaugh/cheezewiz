package system

import (
	"cheezewiz/pkg/ecs"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type RegisterPlayer struct {
	gamepadIDsBuf []ebiten.GamepadID
	gamepadIDs    map[ebiten.GamepadID]ecs.EntityHandle
}

func NewRegisterPlayer() *RegisterPlayer {
	return &RegisterPlayer{}
}

func (rp *RegisterPlayer) Update() {
	if rp.gamepadIDs == nil {
		rp.gamepadIDs = make(map[ebiten.GamepadID]ecs.EntityHandle)
	}

	// Log the gamepad connection events.
	rp.gamepadIDsBuf = inpututil.AppendJustConnectedGamepadIDs(rp.gamepadIDsBuf[:0])
	// for _, id := range rp.gamepadIDsBuf {
	// entry := entity.MakePlayer(world, input.GamePad{ID: id})
	// rp.gamepadIDs[id] = entry
	// }

	// for id, entry := range rp.gamepadIDs {
	// 	if inpututil.IsGamepadJustDisconnected(id) {
	// 		delete(rp.gamepadIDs, id)
	// 		// world.Remove(entry.Entity())
	// 	}
	// }
}
