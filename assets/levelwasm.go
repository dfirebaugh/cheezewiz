//go:build js
// +build js

package assets

import "github.com/lafriks/go-tiled"

// we need to embed the map assets.  However, there is currently an issue preventing us from doing that
//   see: https://github.com/lafriks/go-tiled/issues/63
func GetKitchenMap() (*tiled.Map, error) {
	return &tiled.Map{}, nil
}
