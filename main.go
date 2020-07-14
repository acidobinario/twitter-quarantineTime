package main

import (
	"log"
	"os"
	"twitter-quarantineTime/internal/cmd"
)

func main() {
	app := cmd.App()

	if err := app.Run(os.Args); err != nil {
		log.Printf("[ERROR] command failed, exiting with err %s", err)
	}
}
