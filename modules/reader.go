package modules

import (
  "fmt"
  "io/ioutil"
  "sync"
)

type Reader struct {
  id int
  res *Mutex
  read *int
  write *int
  priority string
  wg *sync.WaitGroup
}

func NewReader(id int, res *Mutex, read, write *int, priority string, wg *sync.WaitGroup) Reader {
  r := Reader{id: id, res: res, read: read, write: write, priority: priority, wg: wg}
  return r
}

func (r Reader) Read() {
  fmt.Printf("Reader %d entering non-critical.\n", r.id)
  if r.priority == "write" {
    for *r.write > 0 {}
  }

  *r.read++
  r.res.RLock()
  fmt.Printf("Reader %d entering critical.\n", r.id)
  _, err := ioutil.ReadFile(r.res.File.Name())
  if err != nil {
    panic(err)
  }
  fmt.Printf("Reader %d leaving critical.\n", r.id)
  r.res.RUnlock()
  *r.read--
  r.wg.Done()
}
