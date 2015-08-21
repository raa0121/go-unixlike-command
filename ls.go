package main

import (
  "fmt"
  "flag"
  "path"
  "io/ioutil"
)

func main() {
  flag.Parse()
  p := path.Clean(flag.Arg(0))
  files, _ := ioutil.ReadDir(p)
  for _, f := range files {
    fmt.Println(f.Name())
  }
}
