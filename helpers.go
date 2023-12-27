package main

import (
	"fmt"
	"os"
	"regexp"
)

func validateName(s string) error {
  match, _ := regexp.MatchString("^[a-zA-Z0-9_]+$", s)
  if !match {
    return fmt.Errorf("Name must contain only alphanumeric characters and underscores")
  }
  return nil
}

func resolvePath(p string) error {
  if _,err := os.Stat(p); os.IsNotExist(err) {
  }
  return nil
}
