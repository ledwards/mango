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
  link, _ := mango.Driver.FindElement(selenium.ByXPATH, "//a[contains(., '" + text + "')]")
  newURLString, _ := link.GetAttribute("href")
  newURL, _ := url.Parse(newURLString)
  link.Click()
  mango.currentURL = *newURL
}

func (mango *Mango) ClickCss(selector string) {
  link, _ := mango.Driver.FindElement(selenium.ByCSSSelector, selector)
  newURLString, _ := link.GetAttribute("href")
  newURL, _ := url.Parse(newURLString)
  link.Click()
  mango.currentURL = *newURL
}

func (mango *Mango) FillInLabel(label string) {
}

func (mango *Mango) FillInCss(selector string) {
}

func (mango *Mango) ChooseLabel(label string) {
}

func (mango *Mango) ChooseCss(css string) {
}

func (mango *Mango) CheckLabel(label string) {
}

func (mango *Mango) CheckCss(css string) {
}

func (mango *Mango) UncheckLabel(label string) {
}

func (mango *Mango) UncheckCss(css string) {
}

func (mango *Mango) AttachFileLabel(label string) {
}

func (mango *Mango) AttachFileCss(css string) {
}

func (mango *Mango) SelectLabel(label string) {
}

func (mango *Mango) SelectCss(css string) {
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

func (mango *Mango) HasContent(content string) bool {
}

func (mango *Mango) HasCss(css string) bool {
  elem, _ := mango.Driver.FindElement(selenium.ByCSSSelector, css)
  if elem == nil {
    return false
  } else {
    return true
  }
}
