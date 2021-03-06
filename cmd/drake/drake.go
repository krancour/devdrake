package main

import (
	"fmt"
	"os"

	"github.com/lovethedrake/devdrake/pkg/version"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "drake"
	app.Usage = "unified, container-aware task execution"
	app.Version = version.Version()
	if version.Commit() != "" {
		app.Version = fmt.Sprintf("%s+%s", app.Version, version.Commit())
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  flagsFile,
			Usage: "specify the location of drake configuration",
			Value: "Drakefile.yaml",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:      "run",
			Usage:     "execute drake jobs(s) or pipeline(s)",
			UsageText: "drake run name... [options]",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  flagsPipeline,
					Usage: "execute a pipeline instead of a job",
				},
				cli.BoolFlag{
					Name:  flagsDebug,
					Usage: "display debug info",
				},
				cli.BoolFlag{
					Name:  flagsConcurrently,
					Usage: "enable concurrent job execution",
				},
				cli.StringFlag{
					Name:  flagsSecretsFile,
					Usage: "specify the location of drake secrets",
					Value: "Drakesecrets",
				},
			},
			Action: run,
		},
	}
	fmt.Println()
	if err := app.Run(os.Args); err != nil {
		fmt.Printf("\n%s\n\n", err)
		os.Exit(1)
	}
	fmt.Println()
}
