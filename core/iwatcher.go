package core

// IWatcher is an interface for different types of
// watchers. At the moment, the only watcher I can
// think of, would be a FileSystemWatcher.
type IWatcher interface {
	WatcherType() WatcherType
	IsWatching() bool
	Watch() bool
	Stop() bool
}

// WatcherType is a string to identity the type of
// watcher, but the real purpose to ensure the
// concrete impelementation only satisfies the
// IWatcher interface
type WatcherType string
