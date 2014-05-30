mango
=====

A Selenium-based integration testing framework for Go

Setup
-----

### OS X:

`brew install selenium-server-standalone`

Make sure selenium-server-standalone is running (run `brew info` for instruction)

Supported Methods
-----------------

### Navigating
* Visit(pathString)
* ClickLinkText
* ClickCss

### Properties
* CurrentPath()
* CurrentURL

### Booleans
* HasCss(cssString)

Tests
-----
Make sure to run `cd fixture; go run main.go` in another window before running the tests. Future versions will run this test server for you in the background when you run the test suite.
