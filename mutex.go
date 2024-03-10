package main

import "sync"

type Foo struct {
  mu sync.Mutex
  value int64
}

func (f *Foo) Incr() {
  f.mu.Lock()
  defer f.mu.Unlock()
  f.value++
}
