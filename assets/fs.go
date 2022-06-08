package assets

import (
	"embed"
	_ "embed"
)

//go:embed scenes/*
var SceneFS embed.FS

//go:embed *
var AssetFS embed.FS
