//go:build !js
// +build !js

package dentity

import (
	"os"

	"github.com/sirupsen/logrus"
)

func pathToBytes(path string) []byte {
	f, err := os.ReadFile(path)
	if err != nil {
		logrus.Fatal(err)
	}
	return f
}
