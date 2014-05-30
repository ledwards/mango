package mango

import (
  "testing"
)

func TestVisitsPage(t *testing.T) {
  mango := Start("http://localhost:3000")
  defer mango.Quit()
  mango.Visit("/")

  if mango.CurrentPath() != "/" {
    t.Error("Expected path to be '/'. Got " + mango.CurrentPath())
  }
}

func TestClickCss(t *testing.T) {
  mango := Start("http://localhost:3000")
  defer mango.Quit()
  mango.Visit("/")
  mango.ClickCss("a#link-test")

  if !mango.HasCss("p#linked-to") {
    t.Error("Expected to be on the \"linked to\" page")
  }

  if mango.CurrentPath() != "/link" {
    t.Error("Expected path to be '/link'. Got " + mango.CurrentPath())
  }
}

func TestClickCssElementDoesNotExist(t *testing.T) {
  // Test should fail rather than blowing up
  //pending
}

func TestClickLinkText(t *testing.T) {
  mango := Start("http://localhost:3000")
  defer mango.Quit()
  mango.Visit("/")
  mango.ClickLinkText("Link Test")

  if !mango.HasCss("p#linked-to") {
    t.Error("Expected to be on the \"linked to\" page")
  }

  if mango.CurrentPath() != "/link" {
    t.Error("Expected path to be '/link'. Got " + mango.CurrentPath())
  }
}

func TestHasCssReturnsTrue(t *testing.T) {
  mango := Start("http://localhost:3000")
  defer mango.Quit()
  mango.Visit("/")

  if !mango.HasCss("body#mango-fixture") {
    t.Error("Expected page to have CSS 'body#mango-fixture'")
  }
}

func TestHasCssReturnsFalse(t *testing.T) {
  mango := Start("http://localhost:3000")
  defer mango.Quit()
  mango.Visit("/")

  if mango.HasCss("#this-id-doesnt-exist") {
    t.Error("Expected page to not have CSS '#google-fail'")
  }
}
