//go:build js
// +build js

package dentity

import (
	"cheezewiz/assets"
	"cheezewiz/config/entities"
)

// for assets/configs to be included in the wasm build,
//   we have to embed them into the binary
//   we may have to look at generating this file

type pathLabel string

const (
	jellyBeanGreen   pathLabel = "./config/entities/jellybeangreen.entity.json"
	jellyBeanPink    pathLabel = "./config/entities/jellybeanpink.entity.json"
	jellyBeanBlue    pathLabel = "./config/entities/jellybeanblue.entity.json"
	jellyBeanRainbow pathLabel = "./config/entities/jellybeanrainbow.entity.json"

	jellyBeanGreenPNG   pathLabel = "./assets/jellybeangreen.png"
	jellyBeanPinkPNG    pathLabel = "./assets/jellybeanpink.png"
	jellyBeanBluePNG    pathLabel = "./assets/jellybeanblue.png"
	jellyBeanRainbowPNG pathLabel = "./assets/jellybeanrainbow.png"
)

var embededLookup = map[pathLabel][]byte{
	jellyBeanGreen:      entities.JellyBeanGreenRaw,
	jellyBeanPink:       entities.JellyBeanPinkRaw,
	jellyBeanBlue:       entities.JellyBeanBlueRaw,
	jellyBeanRainbow:    entities.JellyBeanRainbowRaw,
	jellyBeanGreenPNG:   assets.JellyBeanGreenRaw,
	jellyBeanPinkPNG:    assets.JellyBeanPinkRaw,
	jellyBeanBluePNG:    assets.JellyBeanBlueRaw,
	jellyBeanRainbowPNG: assets.JellyBeanRainbowRaw,
}
