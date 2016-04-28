package main

import (
  "testing"
)

func TestPrintHello(t *testing.T) {
  name := "big world"
  helloWorld := PrintHello(name)
  if helloWorld != "hello, big world!" {
    t.Error("output string should be \"hello, big world!\"")
  }
}
