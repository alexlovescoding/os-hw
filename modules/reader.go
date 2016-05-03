package modules

import "io/ioutil"

type Reader struct {
  id int
  filename string
  read chan bool
  write chan bool
  priority string
}

func NewReader(id int, filename string, read chan bool, write chan bool, priority string) Reader {
  p := Reader{id: id, filename: filename, read: read, write: write, priority: priority}
  return p
}

func (p Reader) Read() {
  p.read <- true
  time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond)
  <- p.read
}
