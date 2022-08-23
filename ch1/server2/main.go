package main

import (
	"fmt"
	"log"
	"net/http"
	""
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		lissajous(m)
	}
	http.HandleFunc("/hello", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}


}