// main.go
package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
    _ "github.com/lib/pq"
)

func main() {
    connStr := "host=db user=postgres password=mysecretpassword dbname=mydb sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello, World!")
    })

    log.Fatal(http.ListenAndServe(":8080", nil))
}
