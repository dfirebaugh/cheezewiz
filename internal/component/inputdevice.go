package component

import (
	"cheezewiz/internal/input"
)

type InputDevice struct {
	Device input.PlayerInput `json:"device"`
}
