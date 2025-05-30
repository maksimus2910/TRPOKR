package main

import (
  "database/sql"
  "log"
  "os"

  _ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
  dsn := os.Getenv("DATABASE_URL")
  if dsn == "" {
    dsn = "postgres://user:password@db:5432/fruitshop?sslmode=disable"
  }
  var err error
  DB, err = sql.Open("postgres", dsn)
  if err != nil {
    log.Fatal(err)
  }
  if err = DB.Ping(); err != nil {
    log.Fatal(err)
  }
}
