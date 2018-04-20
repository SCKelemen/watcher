package core

type IWatcher interface {
	WatcherType() WatcherType
}

type WatcherType string
