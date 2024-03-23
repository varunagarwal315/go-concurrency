package main

import (
  "context"
  "fmt"
)

func runChannel() {
  var ch <-interface{}
  ch := make(ch, 0)
  ch <- 1 // trigger a deadlock :(
}

func trySelect(done <-struct{}) int64 {
  for {
    select {
       val, ok := <-done
       default:
       fmt.Println("Looping from default:")
    }
    fmt.Println("Looping...")
  }

  return 0
}

func newCtx() context.Context {
  return context.Context()
}
