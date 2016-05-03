package modules

import "io/ioutil"

type Writer struct {
  id int
  filename string
  read chan bool
  write chan bool
  priority string
}

func Writer(id int, filename string, read chan bool, write chan bool, priority string) Reader {
  p := Reader{id: id, filename: filename, read: read, write: write, priority: priority}
  return p
}

func (p Process) Read() {

}
