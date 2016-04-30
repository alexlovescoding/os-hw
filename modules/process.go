package modules

type Process struct {
  id int
  data []string
}

func NewProcess(id int) Process {
  p := Process{id: id, data: make([]string, 0)}
  return p
}

func (p *Process) Write(write chan<- []string, s string) {

  p.data = append(p.data, s)
  write <- p.data
}

func (p Process) Read(read chan<- []string) {
  read <- p.data
}
