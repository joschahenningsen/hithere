package main

import (
	"fmt"
	"os"
	"net/http"
	"log"
)

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello From %s\n", os.Getenv("host"))
	})
	log.Fatal(http.ListenAndServe(":8083", nil))
}
