package main

import (
	"html/template"
	"log"
	"net/http"
)
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Define the data for the template
	data := struct {
		Title   string
		Message string
	}{
		Title:   "Hello, Dockerized Go!",
		Message: "This is a Go web app with HTML templating.",
	}

	// Try inline parsing for troubleshooting
	templateContent := `
<!DOCTYPE html>
<html>
<head>
    <title>Hello Go!</title>
    <script>console.log("hello")</script>
</head>
<body>
    <h1>{{.Title}}</h1>
    <p>{{.Message}}</p>
</body>
</html>
`
	tmpl, err := template.New("index").Parse(templateContent)
	if err != nil {
		http.Error(w, "Failed to parse template", http.StatusInternalServerError)
		log.Println("Error:", err)  // Log the error
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Println("Error executing template:", err) // Log template execution error
		return
	}

	log.Println("Success: Served request from", r.RemoteAddr)  // Log a successful request
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":8000", nil)
}
