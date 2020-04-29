package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
	"os/exec"
)

var (
	app      = cli.NewApp()
	endPoint string
	myFlags  = []cli.Flag{
		&cli.StringFlag{
			Name:        "argument",
			Value:       "Enter the Argument",
			Aliases:     []string{"arg"},
			Destination: &endPoint,
		},
	}
)

func handleError(err error) {
	if err != nil {
		log.Fatalln("Error::", err)
	}
}
func info() {
	app.Name = "Utility Cli"
	app.Usage = "A tool for querying Linux Commands"
	app.Version = "0.0.1"
}
func command() {
	app.Commands = []*cli.Command{
		{
			Name:    "list",
			Usage:   "Use to List Values",
			Flags:   myFlags,
			Aliases: []string{"l"},
			Action: func(context *cli.Context) error {
				cmd, err := exec.Command("ls", "-lrth", endPoint).CombinedOutput()
				handleError(err)
				fmt.Println(string(cmd))
				return nil
			},
		},
		{
			Name:    "grep",
			Usage:   "Use to List Values",
			Flags:   myFlags,
			Aliases: []string{"g"},
			Action: func(context *cli.Context) error {
				cmd, err := exec.Command("grep", "-ir", endPoint).CombinedOutput()
				handleError(err)
				fmt.Println(string(cmd))
				return nil
			},
		},
	}
}
func main() {
	info()
	command()
	err := app.Run(os.Args)
	handleError(err)
}
