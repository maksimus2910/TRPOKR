package main

import (
  "log"
  "net/http"
  "os"

  "github.com/gorilla/mux"
)

func main() {
  InitDB()
  defer DB.Close()

  r := mux.NewRouter()

  r.HandleFunc("/products", GetProducts).Methods("GET")
  r.HandleFunc("/products", AddProduct).Methods("POST")
  r.HandleFunc("/products/{id}", DeleteProduct).Methods("DELETE")

  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
  }

  log.Printf("Server started on port %s", port)
  log.Fatal(http.ListenAndServe(":"+port, r))
}

