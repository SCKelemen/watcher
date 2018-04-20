package core

func NewWatcher() Watcher {
	return Watcher{}
}

type Watcher struct {
	watchers   []IWatcher
	isWatching bool
}

func (w Watcher) AddWatcher(watcher IWatcher) bool {
	w.watchers = append(w.watchers, watcher)
	return true
}

// Watch starts the watcher, and
// returns the state of the watcher
func (w Watcher) Watch() bool {
	w.isWatching = true
	return w.isWatching
}

func (w Watcher) Stop() bool {
	w.isWatching = false
	return w.isWatching
}
