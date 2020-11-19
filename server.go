package main

import "github.com/go-martini/martini"

func main() {
  m := martini.Classic()
  m.Get("/", func() string {
    return "Hallo Welt!"
  })
  m.RunOnAddr(":8080")
  m.Run()
}
