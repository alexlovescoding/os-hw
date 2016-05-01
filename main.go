package main

import (
  "fmt"
  "github.com/alexlovescoding/os-hw/modules"
  "strconv"
)

func main() {
  p := modules.NewProcess(1)
  write := make(chan []string)
  go p.Write(write, strconv.Itoa(0))
  for i := 1; i < 20; i++ {
    select {
    case <-write:
      go p.Write(write, strconv.Itoa(i))
    }
  }
  read := make(chan []string)
  select {
  case <-write:
    go p.Read(read)
  }
  fmt.Println(<-read)
}
