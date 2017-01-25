package main

import (
	"github.com/ajstarks/svgo"
	"log"
	"net/http"
)

func main() {
	http.Handle("/circle.svg", http.HandlerFunc(circle))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func circle(w http.ResponseWriter, req *http.Request) {
	// w.Header().Set("Content-Type", "image/svg+xml")
	s := svg.New(w)
	s.Start(500, 500)
	s.Circle(250, 250, 125, "fill:blue;stroke:black")
	s.End()
}
