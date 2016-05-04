package modules

import (
  "fmt"
  "sync"
)

type Writer struct {
  id int
  res *Mutex
  read *int
  write *int
  priority string
  wg *sync.WaitGroup
}

func NewWriter(id int, res *Mutex, read, write *int, priority string, wg *sync.WaitGroup) Writer {
  w := Writer{id: id, res: res, read: read, write: write, priority: priority, wg: wg}
  return w
}

func (w Writer) Write() {
  fmt.Printf("Writer %d entering non-critical.\n", w.id)

  *w.write++
  if w.priority == "read" {
    for *w.read > 0 {}
  }

  w.res.Lock()
  fmt.Printf("Writer %d entering critical.\n", w.id)
  data := fmt.Sprintf("Writer #%d\n", w.id)
  w.res.File.WriteString(data)
  fmt.Printf("Writer %d leaving critical.\n", w.id)
  w.res.Unlock()
  *w.write--
  w.wg.Done()
}
