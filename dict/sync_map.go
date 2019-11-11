package dict

import "sync"

type SyncMap struct {
	m    sync.Map
	size int
}

func NewSyncMap() *SyncMap {
	return &SyncMap{m: sync.Map{}}
}

func (sm *SyncMap) Store(k, v interface{}) {
	_, loaded := sm.m.LoadOrStore(k, v)
	if !loaded {
		sm.size++
	}
}

func (sm *SyncMap) Load(k interface{}) (v interface{}, ok bool) {
	return sm.m.Load(k)
}

func (sm *SyncMap) Contains(k interface{}) bool {
	_, ok := sm.m.Load(k)
	return ok
}

func (sm *SyncMap) Get(k interface{}) interface{} {
	v, _ := sm.m.Load(k)
	return v
}

func (sm *SyncMap) Delete(k interface{}) {
	if _, ok := sm.m.Load(k); ok {
		sm.m.Delete(k)
		sm.size--
	}
}

func (sm *SyncMap) Size() int {
	return sm.size
}
