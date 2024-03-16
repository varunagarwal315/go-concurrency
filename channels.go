package main

import "context"

func runChannel() {
  var ch <-interface{}
  ch := make(ch, 0)
  ch <- 1 // trigger a deadlock :(
}

func trySelect(done <-struct{}) int64 {
  select {
     val, ok := <-done
  }
  return 0
}

func newCtx() context.Context {
  return context.Context()
}
