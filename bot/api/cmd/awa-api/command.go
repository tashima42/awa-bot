package awaapi

import (
	"github.com/spf13/cobra"
	"github.com/tashima42/awa-bot/bot/api"
	"github.com/tashima42/awa-bot/bot/pkg/auth"
	"github.com/tashima42/awa-bot/bot/pkg/db"
	"log"
	"os"
)

func Command(repo *db.Repo) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "api",
		Short: "awa-api is the REST API for Awa bot",
		Long:  "awa-api is the REST API for Awa bot",
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("getting hash helper instance")
			hashHelper := auth.GetHashHelperInstance()
			log.Println("getting jwt helper instance")
			jwtHelper := auth.NewJWTHelper([]byte(os.Getenv("TELEGRAM_TOKEN")))
			log.Println("running awa-api command")
			api.Serve(repo, hashHelper, jwtHelper)
		},
	}
	return rootCmd
}
