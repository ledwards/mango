package fixture

import (
  "github.com/codegangsta/martini"
  "github.com/codegangsta/martini-contrib/render"
)

func main() {
  // Should I use Classic?
  m := martini.New()
  m.Use(render.Renderer())
  r := martini.NewRouter()

  r.Get("/", func(r render.Render) {
    r.HTML(200, "index", nil)
  })

  r.Get("/link", func(r render.Render) {
    r.HTML(200, "link", nil)
  })

  m.Action(r.Handle)
  m.Run()
}

