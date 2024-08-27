package main

import (
	"log"
	"net/http"
	"text/template"
)

var html string = `<html>
<body style="background-color:#0902db">
<h1 style="color:#ffffff">Hello world!</h1>
<ul>
<li style="color:#ffffff">Host: {{.Host}}</li>
<li style="color:#ffffff">URL: {{.URL}}</li>
<li style="color:#ffffff">RequestURI: {{.RequestURI}}</li>
<li style="color:#ffffff">Method: {{.Method}}</li>
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
