package main

import (
  "encoding/json"
  "github.com/go-martini/martini"
  "github.com/martini-contrib/render"
  "log"
  "net/http"
)

type Person struct {
  Name     string    `json:"name"`
}

func main() {
  m := martini.Classic()
  // render html templates from views directory
  m.Use(render.Renderer(render.Options{
    Directory: "views", // Specify what path to load the templates from.
    Extensions: []string{".tmpl", ".html"}, // Specify extensions to load for templates.
    IndentJSON: true, // Output human readable JSON
  }))

  // Routes
  // http://localhost:3000/ serves by default the static page at './public/index.html'

  // http://localhost:3000/hello
  m.Get("/hello", func() string {
    return "Hello world!"
  })
  
  // http://localhost:3000/hello/myname
  m.Get("/hello/:name", func(params martini.Params, res http.ResponseWriter, req *http.Request) (int, string) {
    return 200, "Hello " + params["name"]
  })
  
  // http://localhost:3000/hello?name=myname
  m.Get("/hello_template", func(r render.Render, req *http.Request) {
     r.HTML(200, "hello", req.URL.Query().Get("name"))
  })
  
  // curl -X POST -i -k -u token: "http://localhost:3000/hello.json" -d '{"name": "MyName"}'
  m.Post("/hello.json", func(r render.Render, req *http.Request) {
    decoder := json.NewDecoder(req.Body)
    var person Person
    err := decoder.Decode(&person)
    if err != nil {
      r.JSON(500, map[string]interface{}{"error": "Failed to parse body"})
      return
    }
    log.Println("Person: "+person.Name)
    r.JSON(http.StatusCreated, map[string]interface{}{"hello": person.Name})
  })
  
  // http://localhost:3000/hello.json?name=myname
  m.Get("/hello.json", func(r render.Render, req *http.Request) {
    type Greeting struct {
      Hello     string    `json:"hello"`
    }

    r.JSON(200, Greeting{Hello: req.URL.Query().Get("name")})
  })
  
  m.Run()
}