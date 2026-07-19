package main

import (
	"strings"
	"time"
)

type Router struct {
	routes map[string]HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		routes: make(map[string]HandlerFunc),
	}
}

func (r *Router) Handle(path string, handler HandlerFunc) {
	r.routes[path] = handler
}

func (r *Router) ServeGemini(w ResponseWriter, req *Request) {
	for pattern, handler := range r.routes {
		if strings.HasPrefix(req.Path, pattern) {
			handler(w, req)
			return
		}
	}
	w.WriteHeader(51, "page not found")
	w.WriteString("# 404\n\n the page you are looking for dose not exists")
}

func homeHandler(w ResponseWriter, r *Request) {
	w.WriteHeader(20, "text/gemini")
	content := `# welcome

## navigation
=> /about about this site
=> /hello hello

## outside link
=> gemini://gemini.circumlunar.space/ Gemini offical website

---
*server time: ` + time.Now().Format("2006-01-02 15:04:05") + `*
`
	w.WriteString(content)
}

func aboutHandler(w ResponseWriter, r *Request) {
	w.WriteHeader(20, "text/gemini")
	content := `# 📖 about this site
this is a site completely written in golang, as a learning project to get better understanding of internet connection
=> / pack to homepage
`
	w.WriteString(content)
}

func helloHandler(w ResponseWriter, r *Request) {
	w.WriteHeader(20, "text/gemini")
	content := `# 👋 Hello Gemini!

this is a hello test

=> / back to home
`
	w.WriteString(content)
}
