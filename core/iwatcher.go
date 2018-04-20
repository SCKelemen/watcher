package core

type IWatcher interface {
	WatcherType() WatcherType
	IsWatching() bool
	Watch() bool
	Stop() bool
}

type WatcherType string
