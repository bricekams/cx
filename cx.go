package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bricekams/cx/db"
	"github.com/urfave/cli/v2"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("cx: ")
	db.Init()
	app := &cli.App{
		Name:    "cxdir",
		Version: "v0.1",
		Authors: []*cli.Author{
			{
				Name:  "Brice Kamhoua",
				Email: "kamhoua.k.brice@gmail.com",
			},
		},
		Copyright: "(c) 2023 Brice Kamhoua",
		Usage:     "A simple shortcuts manager",
		// UsageText: "cx [global options] command [command options] [arguments...]",
		Commands: []*cli.Command{

			{
				Name:    "create",
				Aliases: []string{"c"},
				Usage:   "Create a new shortcut",
				Before: func(cCtx *cli.Context) error {
					name := cCtx.Args().First()
					if name == "" {
						log.Fatal("Please provide a shortcut name")
					}
					err := validateName(name)
					if err != nil {
						log.Fatal(err.Error())
					}
					return nil
				},
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "path", Usage: "Specify the path to be saved", Aliases: []string{"p"}},
				},
				Action: func(cCtx *cli.Context) error {
					name := cCtx.Args().First()
					path := cCtx.String("path")
          var err_path error = nil
          if path == "" {
            path, err_path = os.Getwd()
            if err_path != nil {
              log.Fatal("An error occurred while processing the path")
            }
          }
					exists, err_exists := exists(path)
					if err_exists != nil {
						log.Fatal("An error occured while processing the path")
					}
					if !exists {
            log.Println(path)
						log.Fatal("The given path does not exists")
					}
					rootedPath, err_rootedPath := resolvePath(path)
					if err_rootedPath != nil {
						log.Fatal("An error occured while resolving the path")
					}
					err := db.Create(name, rootedPath)
					if err != nil {
						log.Fatal(err.Error())
					}
					log.Printf("Shortcut created successfully. Use `cx %s`\n", name)
					return nil
				},
			},

			{
				Name:    "update",
				Aliases: []string{"u"},
				Usage:   "Update a shortcut path",
        Before: func(cCtx *cli.Context) error {
          name := cCtx.Args().First()
          if name == "" {
            log.Fatal("Please provide the shortcut name")
          }
          return nil
        },
        Action: func(cCtx *cli.Context) error {
          name := cCtx.Args().First()
          path := db.Get(name)
          var newPath string
          fmt.Println("The path for this shortcut is:", path)
          fmt.Print("Enter the new path: ")
          for {
            fmt.Scan(&newPath)
            var err_newPath error
            if newPath == "" {
              newPath, err_newPath = os.Getwd()
              if err_newPath != nil {
                log.Fatal("An error occurred while processing the path")
              }
            }
            exists, err_exists := exists(newPath)
            if err_exists != nil {
              log.Fatal("An error occured while processing the path")
            }
            if exists {
              break
            }
            fmt.Print("The given path does not exists, enter a new one: ")
          }
          err := db.Update(name,newPath)
          if err != nil {
            log.Fatal(err.Error())
          }
          fmt.Printf("Shortcut updated successfully! Use `cx %s`\n",name)
          return nil
        },
      },

			{
				Name:      "rename",
				Aliases:   []string{"r"},
				Usage:     "rename a shortcut",
				UsageText: "cx rename <old_name> <new_name>",
				Before: func(cCtx *cli.Context) error {
					oldName := cCtx.Args().First()
					newName := cCtx.Args().Get(1)
					if oldName == "" || newName == "" {
						log.Fatal("rename expect two args. see --help")
					}
					err := validateName(oldName)
					err_n := validateName(newName)
					if err != nil {
						log.Fatal(err.Error())
					}
					if err_n != nil {
						log.Fatal(err_n.Error())
					}
					return nil
				},
				Action: func(cCtx *cli.Context) error {
					oldName := cCtx.Args().First()
					newName := cCtx.Args().Get(1)
					err := db.Rename(oldName, newName)
					if err != nil {
						log.Fatal(err.Error())
					}
					log.Printf("Shortcut renamed successfully. Use `cx %s`\n",newName)
					return nil
				},
			},

			{
				Name:    "delete",
				Aliases: []string{"d"},
				Usage:   "Delete a shortcut",
        Before: func(cCtx *cli.Context) error {
          name := cCtx.Args().First()
          if name == "" {
            log.Fatal("Please provide the shortcut name")
          }
          return nil
        },
        Flags: []cli.Flag{
          &cli.BoolFlag{Name: "all", Aliases: []string{"a"}, Usage: "Delete all the saved shortcuts"},
        },
				Action: func(cCtx *cli.Context) error {
					name := cCtx.Args().First()
          err := db.Delete(name)
          if err!=nil {
            log.Fatal(err.Error())
          }
          log.Println("Shortcut", name, "has been deleted succesfully")
					return nil
				},
			},

      {
        Name: "list",
        Aliases: []string{"l"},
        Usage: "List all the shortcuts",
        Action: func(ctx *cli.Context) error {
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
