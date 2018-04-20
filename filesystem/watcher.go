package filesystem

import (
	"github.com/sckelemen/watcher/core"
)

func NewWatcher() core.IWatcher {
	return FileSystemWatcher{}
}

type FileSystemWatcher struct {
	WatchedPaths []string
}

func (fsw FileSystemWatcher) WatcherType() core.WatcherType {
	return FileSystemWatcherType
}

const (
	FileSystemWatcherType core.WatcherType = "file_system_watcher"
)
