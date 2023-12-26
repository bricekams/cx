package main

import (
  "fmt"
  "log"
  "os"

  "github.com/urfave/cli/v2"
)

func main () {
  app := &cli.App{
    Name: "cxdir",
    Version: "v0.1",
    Authors: []*cli.Author{
      {
        Name:"Brice Kamhoua",
        Email: "kamhoua.k.brice@gmail.com",
      },
    },
    Copyright: "(c) 2023 Brice Kamhoua",
    Usage: "A simple shortcuts manager",
    // UsageText: "cx [global options] command [command options] [arguments...]",
    Commands: []*cli.Command{
      {
        Name: "create",
        Aliases: []string{"c"},
        Usage: "Create a new shortcut",
        Before: func(cCtx *cli.Context) error {
          name := cCtx.Args().First()
          if name == "" {
            return cli.Exit("Please provide a name",1)
          }
          return validateName(name)
        },
        Flags: []cli.Flag{
          &cli.StringFlag{Name: "path", Value: ".", Usage: "Specify the path to be saved", Aliases: []string{"p"}},
        },
        Action: func(cCtx *cli.Context) error {
         // name:=cCtx.Args().First()

          fmt.Println("shortcut create")
          return nil
        },
      },
      {
        Name: "update",
        Aliases: []string{"u"},
        Usage: "Update a shortcut path",
        Action: func(ctx *cli.Context) error {
          fmt.Println("update shortcut")
          return nil
        },
      },
      {
        Name: "rename",
        Aliases: []string{"r"},
        Usage: "rename a shortcut",
        Action: func(ctx *cli.Context) error {
          fmt.Println("rename shortcut")
          return nil
        },
      },
      {
        Name: "delete",
        Aliases: []string{"d"},
        Usage: "delete a shortcut",
        Flags: []cli.Flag{
          &cli.BoolFlag{Name: "all", Aliases: []string{"a"}, Usage: "Delete all the saved shortcuts"},
        },
        Action: func(ctx *cli.Context) error {
          fmt.Println("delete shortcut")
          return nil
        },
      },
    },
  }
  app.Suggest = true
  if err := app.Run(os.Args); err != nil {
    log.Fatal(err)
  }
}
