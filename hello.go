package main

import (
  "fmt"
)

func PrintHello(name string) string {
  return "Hello, " + name + "!"
}

func main() {
  fmt.Printf(PrintHello("small world"))
}
