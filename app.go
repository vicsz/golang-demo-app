package main

import (
    "fmt"
    "log"
    "net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Helloworld!")

    fmt.Println("Endpoint Hit: Helloworld")
}

func handleRequests() {
    http.HandleFunc("/", helloWorld)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
    handleRequests()
}