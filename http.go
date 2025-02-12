package main

import (
  "os"
  "fmt"
  "net/http"
  "log"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "443"
    }

    fmt.Fprintf(os.Stdout, "Listening on :%s\n", port)
    hostname, _ := os.Hostname()
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(os.Stdout, "I'm %s\n", hostname)
 	fmt.Fprintf(w, "I'm %s\n", hostname)
    })


    log.Fatal(http.ListenAndServeTLS(":" + port,"/run/secrets/server.cert","/run/secrets/server.key", nil))
}

