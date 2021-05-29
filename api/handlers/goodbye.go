package handlers

import (
	"fmt"
	"log"
	"net/http"
)

type Goobbye struct {
	log *log.Logger
}

func NewGoodbye(l *log.Logger) *Goobbye {
	return &Goobbye{l}
}

func (g *Goobbye) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	g.log.Println("handle goodbye request")

	fmt.Fprintf(w, "goodbye")
}
