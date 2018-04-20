package filesystem

import (
	"github.com/sckelemen/watcher/core"
)

// NewWatcher creates a new FileSystemWatcher
// and returns the IWatcher interface
func NewWatcher() core.IWatcher {
	return FileSystemWatcher{}
}

// FileSystemWatcher notifies when a filesystem
// event occurs on a watched resource. It implements
// the IWatcher interface.
type FileSystemWatcher struct {
	WatchedPaths []string
	isWatching   bool
}

// Watch starts the watcher and returns the watch state
func (fsw FileSystemWatcher) Watch() bool {
	fsw.isWatching = true
	return fsw.isWatching
}

// Stop stops the watcher and returns the watch state
func (fsw FileSystemWatcher) Stop() bool {
	fsw.isWatching = false
	return fsw.isWatching
}

// IsWatching returns the state of the watcher
func (fsw FileSystemWatcher) IsWatching() bool {
	return fsw.isWatching
}

// WatcherType defines the type of watcher, and ensures the correct IFace
func (fsw FileSystemWatcher) WatcherType() core.WatcherType {
	return FileSystemWatcherType
}

const (
	// FileSystemWatcherType is the string identitifier for FileSystemWatcher
	FileSystemWatcherType core.WatcherType = "file_system_watcher"
)
