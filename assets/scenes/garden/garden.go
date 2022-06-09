package garden

import _ "embed"

//go:embed garden_with_bounds.tmx
var GardenRaw []byte

//go:embed garden_bounds.tsx
var GardenTSXRaw []byte
