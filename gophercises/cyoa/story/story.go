package cyoa

import (
	"io/ioutil"
	"encoding/json"
	"net/http"
	"html/template"
	"strings"
)

type Story map[string]Chapter

type Chapter struct {
	Title  	 	string   `json:"title"`
	Paragraphs   	[]string `json:"story"`
	Options		[]Option `json:"options"`
}

type Option struct {
	Text	string `json:"text"`
	Chapter	string `json:"arc"`
}

var defaultHandlerTmpl = `

<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <title>Choose Your Own Adventure</title>
</head>
<body>
    <h1>{{.Title}}</h1>
    {{range .Paragraphs}}
    <p>{{.}}</p>
    {{end}}
    <ul>
        {{range .Options}}
        <li><a href="/{{.Chapter}}">{{.Text}}</a></li>
        {{end}}
    </ul>
</body>
</html>`

func init() {
	tpl = template.Must(template.New("").Parse(defaultHandlerTmpl))
}

var tpl *template.Template

type handler struct {
	s Story
	t *template.Template
}

type HandlerOptions func(h *handler)

func WithTemplate(t *template.Template) HandlerOptions {
	return func(h *handler) {
		h.t = t
	}
}

func StoriesHandler (s Story, opts ...HandlerOptions) http.Handler {
	h := handler{s, tpl}
	for _, opt := range opts {
		opt(&h)
	}
	return h
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}
	path = path[1:]
	if chapter, ok := h.s[path]; ok {
		err := h.t.Execute(w, chapter)
		if err != nil {
			http.Error(w, "Something went wrong...", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "Chapter not found.", http.StatusNotFound)

}

func GetStories (filename string, story *Story) (error) {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw, &story)
    	if err != nil {
		return err
	}
	return nil
}