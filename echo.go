package main

import (
  "fmt"
  "flag"
  "strings"
)

func main() {
  flag.Parse()
  fmt.Println(strings.Join(flag.Args(), " "))
}
