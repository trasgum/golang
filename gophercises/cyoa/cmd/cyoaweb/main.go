package main

import (
	"log"
	"fmt"
	"github.com/docopt/docopt-go"
	"github.com/trasgum/gophercises/cyoa/story"
	"net/http"
	"html/template"
	"strings"
)

func main () {
	usage := `Gophercises cyoa.
	Usage:
	cyoaweb [options]

	Examples:
	cyoaweb -h | --help | -v | --version
	cyoaweb --file ./problems.csv

	Options:
	-h --help            show this help message and exit
	-v --version         show version and exit
	-f --file=<file>     a json file [default: gopher.json]
	-p --port=<port>     port to start CYOA http server
`
	arguments, _ := docopt.Parse(usage, nil, true, "0.1.0-SNAPSHOT", true)
	port := arguments["--port"].(string)
	filename := arguments["--file"].(string)

	var story cyoa.Story
	if err := cyoa.GetStories(filename, &story); err != nil {
		panic(err)
	}
	tpl := template.Must(template.New("").Parse(storyTmpl))
	h := cyoa.StoriesHandler(story,
		cyoa.WithTemplate(tpl),
		cyoa.WithPathFunc(pathFn),
	)
	//h := cyoa.StoriesHandler(story)

	mux := http.NewServeMux()
	mux.Handle("/story", h)
	mux.Handle("/", cyoa.StoriesHandler(story))

	fmt.Printf("Starting web server at: %s...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), mux))
}

func pathFn(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	if path == "/story" || path == "/story/" {
		path = "/story/intro"
	}
	return path[len("/story/"):]
}

var storyTmpl = `
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>Choose Your Own Adventure</title>
  </head>
  <body>
    <section class="page">
      <h1>{{.Title}}</h1>
      {{range .Paragraphs}}
        <p>{{.}}</p>
      {{end}}
      <ul>
      {{range .Options}}
        <li><a href="/story/{{.Chapter}}">{{.Text}}</a></li>
      {{end}}
      </ul>
    </section>
    <style>
      body {
        font-family: helvetica, arial;
      }
      h1 {
        text-align:center;
        position:relative;
      }
      .page {
        width: 80%;
        max-width: 500px;
        margin: auto;
        margin-top: 40px;
        margin-bottom: 40px;
        padding: 80px;
        background: #FCF6FC;
        border: 1px solid #eee;
        box-shadow: 0 10px 6px -6px #797;
      }
      ul {
        border-top: 1px dotted #ccc;
        padding: 10px 0 0 0;
        -webkit-padding-start: 0;
      }
      li {
        padding-top: 10px;
      }
      a,
      a:visited {
        text-decoration: underline;
        color: #555;
      }
      a:active,
      a:hover {
        color: #222;
      }
      p {
        text-indent: 1em;
      }
    </style>
  </body>
</html>`
