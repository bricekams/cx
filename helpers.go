package main

import (
  "regexp"

  "github.com/urfave/cli/v2"
)

func validateName(s string) error {
  match, _ := regexp.MatchString("^[a-zA-Z0-9_]+$", s)
  if !match {
    return cli.Exit("Name must contain only alphanumeric characters and underscores",1)
  }
  return nil
}
