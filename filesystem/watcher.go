package filesystem

import (
	"github.com/sckelemen/watcher/core"
)

func NewWatcher() core.IWatcher {
	return FileSystemWatcher{}
}

type FileSystemWatcher struct {
	WatchedPaths []string
	isWatching   bool
}

func (fsw FileSystemWatcher) Watch() bool {
	fsw.isWatching = true
	return fsw.isWatching
}

func (fsw FileSystemWatcher) Stop() bool {
	fsw.isWatching = false
	return fsw.isWatching
}

func (fsw FileSystemWatcher) IsWatching() bool {
	return fsw.isWatching
}

func (fsw FileSystemWatcher) WatcherType() core.WatcherType {
	return FileSystemWatcherType
}

const (
	FileSystemWatcherType core.WatcherType = "file_system_watcher"
)
