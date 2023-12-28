package db

import (
  "log"
  "time"
  "fmt"

  "github.com/boltdb/bolt"
)

var bucketName string = "cx"

func OpenDb() *bolt.DB {
  db, err := bolt.Open("cx.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
  if err != nil {
    log.Fatal(err)
  }
  return db
}

func Init() {
  db := OpenDb()
  db.Update(func(tx *bolt.Tx) error {
    _, err := tx.CreateBucketIfNotExists([]byte(bucketName))
    if err != nil {
      return fmt.Errorf("Error creating bucket: %s", err)
    }
    return nil
  })
  defer db.Close()
  return
}

func Create(name string,path string) error {
  db := OpenDb()
  defer db.Close()
  if Exists(db, name) {
    return fmt.Errorf("A shortcut with that name already exists")
  }
  if PathExists(db,path) {
    return fmt.Errorf("A shortcut already exists for this path")
  }
  err := db.Update(func(tx *bolt.Tx) error {
    b := tx.Bucket([]byte(bucketName))
    err := b.Put([]byte(name),[]byte(path))
    return err
  })

  if err!=nil { return fmt.Errorf("An error occured while creating the shortcut") }
  return nil
}

func Update(name string,newPath string) error {
  db := OpenDb()
  defer db.Close()
  if !Exists(db,name) {
    return fmt.Errorf("A shortcut with that name does not exists")
  }
  err := db.Update(func(tx *bolt.Tx) error {
    b := tx.Bucket([]byte(bucketName))
    err := b.Put([]byte(name),[]byte(newPath))
    return err
  })
  if err!=nil { return fmt.Errorf("An error occured while updating the shortcut path") }
  return nil
}

func Rename(oldName string, newName string) error {
  db := OpenDb()
  defer db.Close()
  if !Exists(db,oldName) {
    return fmt.Errorf("A shortcut with that name does not exists")
  }
  db.Update(func(tx *bolt.Tx) error {
    b := tx.Bucket([]byte(bucketName))
    path := Get(oldName)
    err := b.Put([]byte(newName),path)
    if err!=nil { return fmt.Errorf("An error occured while renaming the shortcut path") }
    err_delete := b.Delete([]byte(oldName))
    if err_delete!=nil { return fmt.Errorf("An error occured while renaming the shortcut") }
    return nil
  })
  return nil
}

func Delete(name string) error {
  db := OpenDb()
  defer db.Close()
  if !Exists(db,name) {
    return fmt.Errorf("A shortcut with that name does not exists")
  }
  err := db.Update(func(tx *bolt.Tx) error {
    b := tx.Bucket([]byte(bucketName))
    err := b.Delete([]byte(name))
    if err!=nil { return fmt.Errorf("An error occured while deleting the shortcut") }
    return nil
  })
  if err!=nil { return fmt.Errorf("An error occured while renaming the shortcut path") }
  return nil
}

func Exists(db *bolt.DB, name string) bool {
  var shortcut []byte
  db.View(func(tx *bolt.Tx) error {
    b := tx.Bucket([]byte(bucketName))
    shortcut = b.Get([]byte(name))
    return nil
  })

  if shortcut != nil { return true }
  return false
}

func PathExists(db *bolt.DB, path string) bool {
  var exists bool = false
  db.View(func(tx *bolt.Tx) error {
    b := tx.Bucket([]byte(bucketName))
    c := b.Cursor()
    for k, v := c.First(); k != nil; k, v = c.Next() {
      if string([]byte(path)) == string(v) {
        exists = true
      }
    }
    return nil
  })
  return exists
}

func Get(name string) []byte {
  db := OpenDb()
  defer db.Close()
  var path []byte
  db.View(func(tx *bolt.Tx) error {
    b := tx.Bucket([]byte(bucketName))
    path = b.Get([]byte(name))
    return nil
  })
  return path
}
