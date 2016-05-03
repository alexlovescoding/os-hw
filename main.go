package main

import (
  "fmt"
  "github.com/alexlovescoding/os-hw/modules"
  "os"
  "path/filepath"
  "sync"
)

func check(err error) {
  if err != nil {
    panic(err)
  }
}

func main() {
  err := os.MkdirAll("output", 0777)
  check(err)

  var rN int
  var wN int

  fmt.Print("Enter number of readers: ")
  fmt.Scanf("%d", &rN)
  fmt.Print("Enter number of writers: ")
  fmt.Scanf("%d", &wN)

  read := 0
  write := 0

  var priority string
  optprio := 0
  fmt.Println("Priority goes to: ")
  fmt.Println("1: Read")
  fmt.Println("2: Write")
  fmt.Scanf("%d", &optprio)

  if optprio == 1 {
    priority = "read"
  } else if optprio == 2 {
    priority = "write"
  } else {
    panic("Please enter either 1 or 2")
  }

  var filename string
  fmt.Print("Enter filename: ")
  fmt.Scanf("%s", &filename)
  filename = filepath.Join("output", filename)

  f, err := os.Create(filename)
  check(err)

  res := modules.NewMutex(f)
  readers := make([]modules.Reader, rN)
  writers := make([]modules.Writer, wN)

  var wg sync.WaitGroup
  wg.Add(rN+wN)
  for i := 0; i < rN; i++ {
    readers[i] = modules.NewReader(i+1, res, &read, &write, priority, &wg)
    go readers[i].Read()
  }

  for i := 0; i < wN; i++ {
    writers[i] = modules.NewWriter(i+1, res, &read, &write, priority, &wg)
    go writers[i].Write()
  }
  wg.Wait()
}
