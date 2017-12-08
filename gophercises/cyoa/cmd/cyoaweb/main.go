package main

import (
	"log"
	"fmt"
	"github.com/docopt/docopt-go"
	"github.com/trasgum/gophercises/cyoa/story"
	"net/http"
	//"html/template"
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
	//tpl := template.Must(template.New("").Parse("Hello!"))
	//h := cyoa.StoriesHandler(story, cyoa.WithTemplate(tpl))
	h := cyoa.StoriesHandler(story)
	fmt.Printf("Starting web server at: %s...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), h))
}

