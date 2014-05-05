package mango

import (
  "bitbucket.org/tebeka/selenium"
  "net/url"
)

type Mango struct {
  Driver selenium.WebDriver
  rootURL url.URL
  currentURL url.URL
}

func Start(root string) Mango {
  caps := selenium.Capabilities{"browserName": "firefox"}
  wd, _ := selenium.NewRemote(caps, "")
  rootURL, _ := url.Parse(root)
  mango := Mango{Driver: wd, currentURL: *rootURL, rootURL: *rootURL}
  return mango
}

func (mango *Mango) Quit() {
  mango.Driver.Quit()
}

// Commands

func (mango *Mango) Visit(path string) {
  newURL, _ := url.Parse(mango.rootURL.String() + path)
  mango.currentURL = *newURL
  mango.Driver.Get(newURL.String())
}

func (mango *Mango) ClickButtonText(text string) {
}

func (mango *Mango) ClickLinkText(text string) {
}

func (mango *Mango) ClickCss(selector string) {
  link, _ := mango.Driver.FindElement(selenium.ByCSSSelector, selector)
  newURLString, _ := link.GetAttribute("href")
  newURL, _ := url.Parse(newURLString)
  link.Click()
  mango.currentURL = *newURL
}

func (mango *Mango) FillInLabel(path string) {
}

func (mango *Mango) FillInCss(path string) {
}

func (mango *Mango) ChooseLabel(path string) {
}

func (mango *Mango) ChooseCss(path string) {
}

func (mango *Mango) CheckLabel(path string) {
}

func (mango *Mango) CheckCss(path string) {
}

func (mango *Mango) UncheckLabel(path string) {
}

func (mango *Mango) UncheckCss(path string) {
}

func (mango *Mango) AttachFileLabel(path string) {
}

func (mango *Mango) AttachFileCss(path string) {
}

func (mango *Mango) SelectLabel(path string) {
}

func (mango *Mango) SelectCss(path string) {
}

func (mango *Mango) Wait(time int) {
}

// Accessors

func (mango *Mango) CurrentPath() string {
  return mango.currentURL.Path
}

func (mango *Mango) CurrentURL() string {
  return mango.currentURL.String()
}

// Booleans

func (mango *Mango) HasContent(path string) {
}

func (mango *Mango) HasCss(css string) bool {
  elem, _ := mango.Driver.FindElement(selenium.ByCSSSelector, css)
  if elem == nil {
    return false
  } else {
    return true
  }
}
