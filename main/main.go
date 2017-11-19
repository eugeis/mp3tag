package main

import (
	"github.com/urfave/cli"
	"os"
	"github.com/eugeis/mp3tag"
)

func main() {

	app := cli.NewApp()
	app.Usage = "Mp3Tag"
	app.Version = "1.0"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "separator, s",
			Usage: "Separator for split index/prefix of file name",
			Value: "-",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "fileNameToTitle",
			Usage: "Set file name to title",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "folder, f",
					Usage: "Path of folder with mp3 files",
				},
			},
			Action: func(c *cli.Context) (err error) {
				mp3tag.FileNameToTitle(c.String("folder"))
				return
			},
		},
		{
			Name:  "fileNamePrefixToTitle",
			Usage: "Set file index to title",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "folder, f",
					Usage: "Path of folder with mp3 files",
				},
			},
			Action: func(c *cli.Context) (err error) {
				mp3tag.FileNamePrefixToTitle(c.String("folder"), c.GlobalString("separator"))
				return
			},
		}, {
			Name:  "fileNamePrefixToFileName",
			Usage: "Set file index to file name",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "folder, f",
					Usage: "Path of folder with mp3 files",
				},
			},
			Action: func(c *cli.Context) (err error) {
				mp3tag.FileNamePrefixToFileName(c.String("folder"), c.GlobalString("separator"))
				return
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		mp3tag.Log.Err("%v", err)
	}
}
