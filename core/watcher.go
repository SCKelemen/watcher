package core

func NewWatcher() Watcher {
	return Watcher{}
}

type Watcher struct {
	watchers   []IWatcher
	IsWatching bool
}

func (w Watcher) AddWatcher(watcher IWatcher) bool {
	w.watchers = append(w.watchers, watcher)
	return true
}

// Watch starts the watcher, and
// returns the state of the watcher
func (w Watcher) Watch() bool {
	w.IsWatching = true
	for _, watcher := range w.watchers {
		watcher.Watch()
	}
	return w.IsWatching
}

func (w Watcher) Stop() bool {
	w.IsWatching = false
	for _, watcher := range w.watchers {
		watcher.Stop()
	}
	return w.IsWatching
}
