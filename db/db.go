package main

import (
	"log"
	"time"

	"github.com/boltdb/bolt"
)

func init() {
  db, err := bolt.Open("cx.db",0600, &bolt.Options{Timeout: 1 * time.Second})
  if err != nil {
    log.Fatal(err)
  }
  defer db.Close()
}
