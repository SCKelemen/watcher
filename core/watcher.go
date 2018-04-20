package core

// NewWatcher returns a new Watcher
func NewWatcher() Watcher {
	return Watcher{}
}

// Watcher is a type that holds other watchers
// it is not a base class
type Watcher struct {
	watchers   []IWatcher
	IsWatching bool
}

// AddWatcher adds an IWatcher to the Watcher container class
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

// Stop stops the watcher, and
// returns the state of the watcher
func (w Watcher) Stop() bool {
	w.IsWatching = false
	for _, watcher := range w.watchers {
		watcher.Stop()
	}
	return w.IsWatching
}
