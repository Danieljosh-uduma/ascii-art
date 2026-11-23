package server

import (
	"ascii-web/ascii"
	"html/template"
	"net/http"
)

// Parse the template
var Temp *template.Template

func init() {
	var err error
	Temp, err = template.ParseFiles("index.html")
	if err != nil {
		// Fallback for tests running in the server directory
		Temp = template.Must(template.ParseFiles("../index.html"))
	}
}

// ascii art type
type Art struct {
	Success bool
	Ascii   string
}

func homeHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	data := Art{
		Success: false,
	}

	err := Temp.Execute(w, data)
	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)
	}
}

func asciiHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	err := r.ParseForm()
	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	// ascii function call
	switch banner {
	case "standard":
		banner = "banner/standard.txt"
	case "shadow":
		banner = "banner/shadow.txt"
	case "thinkertoy":
		banner = "banner/thinkertoy.txt"
	default:
		banner = "banner/standard.txt"
	}

	asciiArt, err := ascii.TransformText(text, banner)
	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)
		return
	}

	data := Art{
		Success: true,
		Ascii:   asciiArt,
	}
	err = Temp.Execute(w, data)
	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)
	}
}
