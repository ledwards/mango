package main

import (
  "testing"
  "github.com/ledwards/mango"
)

func TestVisitsPage(t *testing.T) {
  mango := mango.Start("http://www.google.com")
  defer mango.Quit()
  mango.Visit("/")

  if mango.CurrentPath() != "/" {
    t.Error("Expected path to be '/'. Got " + mango.CurrentPath())
  }
}

func TestHasCssReturnsTrue(t *testing.T) {
  mango := mango.Start("http://www.google.com")
  defer mango.Quit()
  mango.Visit("/")

  if !mango.HasCss("body") {
    t.Error("Expected page to have CSS 'body'")
  }
}

func TestHasCssReturnsFalse(t *testing.T) {
  mango := mango.Start("http://www.google.com")
  defer mango.Quit()
  mango.Visit("/")

  if mango.HasCss("#google-fail") {
    t.Error("Expected page to not have CSS '#google-fail'")
  }
}
