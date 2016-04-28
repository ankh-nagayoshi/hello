package main

import (
  "fmt"
)

func PrintHello(name string) string {
  return "hello, " + name + "!"
}

func main() {
  fmt.Printf(PrintHello("small world"))
}
