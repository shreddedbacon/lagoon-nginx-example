package main

import (
	"html/template"
	"log"
	"net/http"
)

var html string = `<html>
<body style="background-color:#d1d10a">
<h1 style="color:#000000">Hello world!</h1>
<ul>
<li style="color:#000000">Host: {{.Host}}</li>
<li style="color:#000000">URL: {{.URL}}</li>
<li style="color:#000000">RequestURI: {{.RequestURI}}</li>
<li style="color:#000000">Method: {{.Method}}</li>
<ul>
</body>
</html>`

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.New("main").Parse(html))
		tmpl.Execute(w, r)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
