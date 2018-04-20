package main

import (
	"github.com/sckelemen/watcher/core"
	"github.com/sckelemen/watcher/filesystem"
)

func main() {
	watcher := core.NewWatcher()
	fsw := filesystem.NewWatcher()
	watcher.AddWatcher(fsw)
	watcher.Watch()
}
