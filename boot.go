package main

import (
        "net/http"
        "log"
)

func main() {
        http.Handle("/", http.FileServer(http.Dir("assets")))
        if err := http.ListenAndServe(":8008", nil); err != nil {
                log.Fatal("ListenAndServe: ", err)
        }
}
