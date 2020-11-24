package server

import (
	"fmt"

	"github.com/d3ta-go/ms-email-graphql-api/interface/http-apps/graphql/echo"
	"github.com/spf13/cobra"
)

// graphQLCmd represents the graphQL API server command
var graphQLAPICmd = &cobra.Command{
	Use:   "graphqlapi",
	Short: "Shows the graphqlapi server command.",
	Long:  `Shows the graphqlapi server command.`,
	Run: func(cmd *cobra.Command, args []string) {

		if err := echo.StartGraphQLAPIServer(); err != nil {
			fmt.Println("Error while running `StartGraphQLAPIServer()`: ")
			panic(err)
		}
	},
}

func init() {
	ServerCmd.AddCommand(graphQLAPICmd)
}
