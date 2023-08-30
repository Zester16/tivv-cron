package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"stockpull/network"

	"github.com/gorilla/mux"
)

func main() {
	//fmt.Printf("Hello fom golang")
	r := mux.NewRouter()

	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)

	})

	r.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		sb := network.GetAsiaIndex()
		io.WriteString(w, sb)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)

	// if err := http.ListenAndServe(":"+port, nil); err != nil {
	// 	log.Fatal(err)
	// }
	err := http.ListenAndServe(":"+port, r)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)

	}

}
