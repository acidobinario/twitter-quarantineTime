/*
Package update contains the update command.
It measures how many days you've been on quarantine and changes your twitter username
*/
package update

import (
	"fmt"
	"log"
	"time"

	"github.com/acidobinario/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/urfave/cli"
)

// Command is the update command for the CLI
var Command = cli.Command{
	Name:  "update",
	Usage: "updates the twitter username",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:   "access-token",
			EnvVar: "TWITTER_ACCESS_TOKEN",
			Usage:  "twitter app access token",
		},
		cli.StringFlag{
			Name:   "access-secret",
			EnvVar: "TWITTER_ACCESS_TOKEN_SECRET",
			Usage:  "twitter app access token secret",
		},
		cli.StringFlag{
			Name:   "consumer-key",
			EnvVar: "TWITTER_CONSUMER_KEY",
			Usage:  "twitter app consumer key",
		},
		cli.StringFlag{
			Name:   "consumer-secret",
			EnvVar: "TWITTER_CONSUMER_SECRET",
			Usage:  "twitter app consumer secret",
		},
		cli.StringFlag{
			Name:   "date",
			EnvVar: "Q_DATE",
			Usage:  "date from where you started the quarantine, ex 03/13/2020",
		},
		cli.StringFlag{
			Name:   "username-prefix",
			EnvVar: "TWITTER_USER_PREFIX",
			Usage:  "the prefix to the username number of days in quarantine",
		},
	},
	Action: Run,
}

// Run will connect to a DB and run a query until stopped by SIGINT
func Run(cc *cli.Context) error {
	quarantineDate, err := time.Parse("01/02/2006", cc.String("date"))
	if err != nil {
		return fmt.Errorf("[ERROR] Could not parse the date from flag or env, err: %s", err.Error())
	}

	config := oauth1.NewConfig(cc.String("consumer-key"), cc.String("consumer-secret"))
	token := oauth1.NewToken(cc.String("access-token"), cc.String("access-secret"))
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	for {
		now := time.Now()
		days := now.Sub(quarantineDate).Hours() / 24
		newName := fmt.Sprintf("%s stayed at home for %.2f days", cc.String("username-prefix"), days)
		_, _, err = client.Accounts.UpdateProfile(&twitter.UpdateProfileParams{Name: newName})

		if err != nil {
			return fmt.Errorf("[ERROR] Could not update the profile, err: %s", err.Error())
		}

		log.Printf("[DEBUG] Username Updated with: %s\n", newName)
		time.Sleep(time.Duration(30) * time.Minute)
	}
}
