package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	log *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.log.Println("Handle hello request")

	src, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.log.Println("Unable to read request body: ", err)
		http.Error(w, "cannot read request body", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Hello %s", src)
}
