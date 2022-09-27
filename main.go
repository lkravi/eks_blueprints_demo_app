package main

import (
    "fmt"
    "net/http"
	quote "rsc.io/quote/v3"
)

func main() {

    http.HandleFunc("/", hello)
    http.HandleFunc("/proverb", proverb)
	http.HandleFunc("/version", version)


    http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, req *http.Request) {

    fmt.Fprintf(w, "hello visitor \n")
}

func proverb(w http.ResponseWriter, req *http.Request) {

    fmt.Fprintf(w, quote.GoV3())
}

func version(w http.ResponseWriter, req *http.Request) {

    fmt.Fprintf(w, "App version : 1.0\n")
}

