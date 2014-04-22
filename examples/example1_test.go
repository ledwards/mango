package main

import (
  "testing"
  "github.com/ledwards/mango"
)

func TestVisitsPage(t *testing.T) {
  mango := mango.Start()
  defer mango.Quit()
  mango.Visit("/")
  
  if mango.CurrentPath() != "/" {
    t.Error("Expected path to be '/'. Got " + mango.CurrentPath())
  }
}

