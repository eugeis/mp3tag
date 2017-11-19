package main

import (
	"github.com/urfave/cli"
	"os"
)

func main() {

	app := cli.NewApp()
	app.Usage = "Mp3Tag"
	app.Version = "1.0"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "folder, f",
			Usage: "Path of folder with mp3 files",
			Value: ".",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "fileNameToTitle",
			Usage: "Set file name to title",
			Action: func(c *cli.Context) (err error) {
				FileNameToTitle(paramFolder(c))
				return
			},
		}, {
			Name:  "fileNamePrefixToTitle",
			Usage: "Set file name prefix, separated by separator, to title",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "separator, s",
					Usage: "Separator for split index/prefix of file name",
					Value: "-",
				},
			},
			Action: func(c *cli.Context) (err error) {
				FileNamePrefixToTitle(paramFolder(c), paramSeparator(c))
				return
			},
		}, {
			Name:  "fileNamePrefixToFileName",
			Usage: "Set file name prefix, separated by separator, to file name",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "separator, s",
					Usage: "Separator for split index/prefix of file name",
					Value: "-",
				},
			},
			Action: func(c *cli.Context) (err error) {
				FileNamePrefixToFileName(paramFolder(c), paramSeparator(c))
				return
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		Log.Err("%v", err)
	}
}
func paramSeparator(c *cli.Context) string {
	return c.String("separator")
}
func paramFolder(c *cli.Context) string {
	return c.GlobalString("folder")
}
