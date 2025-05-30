package main

import (
  "encoding/json"
  "net/http"
  "strconv"

  "github.com/gorilla/mux"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
  rows, err := DB.Query("SELECT id, category, name, quantity, price FROM products")
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  defer rows.Close()

  products := []Product{}
  for rows.Next() {
    var p Product
    if err := rows.Scan(&p.ID, &p.Category, &p.Name, &p.Quantity, &p.Price); err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
    products = append(products, p)
  }

  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(products)
}

func AddProduct(w http.ResponseWriter, r *http.Request) {
  var p Product
  if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
    http.Error(w, "Invalid request payload", http.StatusBadRequest)
    return
  }

  err := DB.QueryRow(
    "INSERT INTO products (category, name, quantity, price) VALUES ($1, $2, $3, $4) RETURNING id",
    p.Category, p.Name, p.Quantity, p.Price,
  ).Scan(&p.ID)

  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(p)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  idStr := vars["id"]
  id, err := strconv.Atoi(idStr)
  if err != nil {
    http.Error(w, "Invalid ID", http.StatusBadRequest)
    return
  }

  _, err = DB.Exec("DELETE FROM products WHERE id=$1", id)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.WriteHeader(http.StatusNoContent)
}
