package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadFile("./html/index.html")
		fmt.Fprintf(w, string(body))
	})
	log.Fatal(http.ListenAndServe(":"+os.Args[1], nil))
}
