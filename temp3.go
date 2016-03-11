package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	
	mux.Handle("/",&myHandler{})
	mux.HandleFunc("/hello",sayHello{})
	
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}

type myHandler struct{}

func (*myHandler) ServerHTTP（w http.ResponseWriter, r *http.Request）{
	io.WriteString(w,"URL:"+r.URL.string())
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world , this is version 2.")
}
