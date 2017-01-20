package main

import (
        "fmt"
        "log"
        "net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello World!")
}


func main() {
        http.HandleFunc("/", helloWorld)
        fmt.Println("Started listening...")
        err := http.ListenAndServe(":5555", nil)
        if err != nil {
                log.Fatal("ListenAndServe: ", err)
        }
}
