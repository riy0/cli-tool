package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"

	"github.com/urfave/cli/v2"
)

const (
	todoFilename = ".todo"
)

func getStorageFile() string {
	filename := ""
	existNowTodo := false
	nowDir, err := os.Getwd()
	if err == nil {
		filename = filepath.Join(nowDir, todoFilename)
		_, err = os.Stat(filename)
		if err == nil {
			existNowTodo = true
		}
	}

	if !existNowTodo {
		home := os.Getenv("HOME")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		filename = filepath.Join(home, todoFilename)
	}
	return filename
}

func main() {
	filename := getStorageFile()

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "list",
				Usage:   "task on a list",
				Aliases: []string{"l"},
				Action: func(c *cli.Context) error {
					fmt.Println(filename)
					list(filename)
					return nil
				},
			},
			{
				Name:    "add",
				Usage:   "add a task",
				Aliases: []string{"a"},
				Action: func(c *cli.Context) error {
					add(filename, c.Args().Get(0))
					return nil
				},
			},
			{
				Name:    "done",
				Usage:   "done tasks",
				Aliases: []string{"d"},
				Action: func(c *cli.Context) error {
					done(filename, c.Args().Slice())
					return nil
				},
			},
			{
				Name:    "undone",
				Usage:   "undone tasks",
				Aliases: []string{"u"},
				Action: func(c *cli.Context) error {
					undone(filename, c.Args().Slice())
					return nil
				},
			},
			{
				Name:    "delete",
				Usage:   "delete a task",
				Aliases: []string{"r"},
				Action: func(c *cli.Context) error {
					delete(filename, c.Args().Slice())
					return nil
				},
			},
			{
				Name:    "clear",
				Usage:   "clear done tasks",
				Aliases: []string{"c"},
				Action: func(c *cli.Context) error {
					clear(filename)
					list(filename)
					return nil
				},
			},
			{
				Name:    "sort",
				Usage:   "sort tasks",
				Aliases: []string{"s"},
				Action: func(c *cli.Context) error {
					sortTasks(filename)
					list(filename)
					return nil
				},
			},
			{
				Name:    "rename",
				Usage:   "rename a task",
				Aliases: []string{"re"},
				Action: func(c *cli.Context) error {
					id, err := strconv.Atoi(c.Args().Get(0))
					if err != nil {
						return err
					}
					rename(filename, c.Args().Get(1), id)
					list(filename)
					return nil
				},
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
