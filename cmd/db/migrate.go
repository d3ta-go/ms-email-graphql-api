package db

import (
	"fmt"

	"github.com/d3ta-go/ms-email-graphql-api/interface/cmd-apps/database"
	"github.com/spf13/cobra"
)

// migrateCmd represents the db migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Shows the db migrate command.",
	Long:  `Shows the db migrate command.`,
	Run: func(cmd *cobra.Command, args []string) {

		if err := database.RunDBMigrate(); err != nil {
			fmt.Println("Error while running `RunDBMigrate()`: ")
			panic(err)
		}
	},
}

func init() {
	DBCmd.AddCommand(migrateCmd)
}
