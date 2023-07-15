package build

import (
	"runtime/debug"
)

var Version = "DEV"

func init() {
	if Version == "DEV" {
		if info, ok := debug.ReadBuildInfo(); ok && info.Main.Version != "" {
			Version = info.Main.Version
		}
	}
}
