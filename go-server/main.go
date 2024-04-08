package main

import (
	"fmt"
	"log"
	"net/http"
)

func hellohandle(w http.ResponseWriter, r *http.Request) {
	fmt.Printf(r.Host)
	fmt.Println("hello handle")
}

func funhandle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("fun handle")
}
func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", hellohandle)
	http.HandleFunc("/fun", funhandle)

	fmt.Printf("Server started at port 8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
