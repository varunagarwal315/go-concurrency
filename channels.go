package main

func runChannel() {
  var ch <-interface{}
  ch := make(ch, 0)
  ch <- 1 // trigger a deadlock :(
}
