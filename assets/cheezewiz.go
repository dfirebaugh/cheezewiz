package assets

import (
	_ "embed"
)

//go:embed cheezewiz.png
var CheezeWizRaw []byte

//go:embed cheezewiz-damaged.png
var CheezeWizHurtRaw []byte

//go:embed cheezewiz.slot.png
var CheezeWizSlotRaw []byte
