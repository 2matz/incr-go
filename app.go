package main

import (
    "fmt"
    "net/http"
    "strings"
    "log"
)

func incrFunc(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "incr") //ここでwに入るものがクライアントに出力されます。
}

func countFunc(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "count") //ここでwに入るものがクライアントに出力されます。
}

func main() {
    http.HandleFunc("/incr", incrFunc) //アクセスのルーティングを設定します。
    http.HandleFunc("/count", countFunc) //アクセスのルーティングを設定します。
    err := http.ListenAndServe(":9090", nil) //監視するポートを設定します。
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
