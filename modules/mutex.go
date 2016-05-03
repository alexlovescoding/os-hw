package modules

import (
  "sync"
  "os"
)

type Mutex struct {
  sync.RWMutex
  File *os.File
}

func NewMutex(f *os.File) *Mutex {
  return &Mutex{File:f}
}
