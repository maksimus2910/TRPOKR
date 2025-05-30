package main

import (
  "net/http"
  "net/http/httptest"
  "strings"
  "testing"
)

func TestGetProducts(t *testing.T) {
  req, _ := http.NewRequest("GET", "/products", nil)
  rr := httptest.NewRecorder()
  handler := http.HandlerFunc(GetProducts)

  handler.ServeHTTP(rr, req)

  if status := rr.Code; status != http.StatusOK {
    t.Errorf("expected status %v, got %v", http.StatusOK, status)
  }
}

func TestAddProduct(t *testing.T) {
  payload := `{"category":"Фрукт","name":"Яблоко","quantity":10,"price":25.5}`
  req, _ := http.NewRequest("POST", "/products", strings.NewReader(payload))
  rr := httptest.NewRecorder()
  handler := http.HandlerFunc(AddProduct)

  handler.ServeHTTP(rr, req)

  if status := rr.Code; status != http.StatusOK {
    t.Errorf("expected status %v, got %v", http.StatusOK, status)
  }
}

func TestDeleteProduct(t *testing.T) {
  req, _ := http.NewRequest("DELETE", "/products/1", nil)
  rr := httptest.NewRecorder()
  handler := http.HandlerFunc(DeleteProduct)

  handler.ServeHTTP(rr, req)

  if status := rr.Code; status != http.StatusNoContent {
    t.Errorf("expected status %v, got %v", http.StatusNoContent, status)
  }
}
