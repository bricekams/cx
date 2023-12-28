package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

func validateName(s string) error {
  match, _ := regexp.MatchString("^[a-zA-Z0-9_]+$", s)
  if !match {
    return fmt.Errorf("Name must contain only alphanumeric characters and underscores")
  }
  return nil
}

func resolvePath(p string) (string,error) {
  rootedPath,err := filepath.Abs(p)
  if err != nil {
    return "", fmt.Errorf("Error while resolving the rooted path")
  }
  return rootedPath, nil
}

func exists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil { return true, nil }
    if os.IsNotExist(err) { return false, nil }
    return false, err
}
