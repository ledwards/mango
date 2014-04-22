package mango

import (
  "bitbucket.org/tebeka/selenium"
)

type Mango struct {
  Driver selenium.WebDriver
  currentPath string
}

func Start() Mango {
  caps := selenium.Capabilities{"browserName": "firefox"}
  wd, _ := selenium.NewRemote(caps, "")
  mango := Mango{Driver: wd, currentPath: ""}
  return mango
}

func (mango *Mango) Visit(path string) {
  mango.Driver.Get("127.0.0.1:3000" + path)
  mango.currentPath = path
}

func (mango *Mango) Quit() {
  mango.Driver.Quit()
}

func (mango *Mango) CurrentPath() string {
  return mango.currentPath
}
