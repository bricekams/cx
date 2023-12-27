package main

import (
  "fmt"
  "log"
  "os"

  "github.com/urfave/cli/v2"
  "github.com/bricekams/cx/db"
)

func main () {
  log.SetFlags(0)
  log.SetPrefix("cx: ")
  db.Init()
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
            log.Fatal("Please provide a shortcut name")
          }
          err := validateName(name)
          if err!=nil {
            log.Fatal(err.Error())
          }
          return nil
        },
        Flags: []cli.Flag{
          &cli.StringFlag{Name: "path",Usage: "Specify the path to be saved", Aliases: []string{"p"}},
        },
        Action: func(cCtx *cli.Context) error {
          name := cCtx.Args().First()
          path := cCtx.String("path")
          fmt.Printf("the path is %s\n",path)
          if path == "" {
            path,_ = os.Getwd()
          }
          if _, err := os.Stat(path); os.IsNotExist(err) {
            fmt.Println(path, "does not exist")
          } else {
            fmt.Println("The provided directory named", path, "exists")
          }
          err := db.Create(name,path)
          if err != nil {
            log.Fatal(err.Error())
          }
          log.Printf("Shortcut created successfully. Use `cx %s`\n",name)
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
        UsageText: "cx rename <old_name> <new_name>",
        Before: func(cCtx *cli.Context) error {
          oldName := cCtx.Args().First()
          newName := cCtx.Args().Get(1)

          if oldName == "" || newName == "" {
            log.Fatal("rename expect two args. see --help")
          }
          err := validateName(oldName)
          err_n := validateName(newName)
          if err!=nil {
            log.Fatal(err.Error())
          }
          if err_n!=nil {
            log.Fatal(err_n.Error())
          }
          return nil
        },
        Action: func(ctx *cli.Context) error {
          err := db.Rename(ctx.Args().First(),ctx.Args().Get(1))
          if err != nil {
            log.Fatal(err.Error())
          }
          log.Printf("Shortcut renamed successfully.\n Use `cx ctx.Args().Get(1)`")
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
