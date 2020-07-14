package cmd

import (
	"log"
	"os"
	"twitter-quarantineTime/internal/cmd/update"

	"github.com/hashicorp/logutils"
	"github.com/urfave/cli"
)

func App() *cli.App {
	app := cli.NewApp()
	app.Name = "twitter-quarantineTime"
	app.Usage = "Changes the username with the amount of time you've been on quarantine"
	//app.Version = version.Description()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "log-level",
			EnvVar: "LOG_LEVEL",
			Value:  "DEBUG",
		},
	}
	app.Before = enableLogger
	app.Commands = []cli.Command{update.Command}
	return app
}

func enableLogger(cc *cli.Context) error {
	filter := &logutils.LevelFilter{
		Levels:   []logutils.LogLevel{"TRACE", "DEBUG", "INFO", "WARN", "ERROR"},
		MinLevel: logutils.LogLevel(cc.GlobalString("log-level")),
		Writer:   os.Stdout,
	}
	log.SetOutput(filter)

	return nil
}
