package cli

import (
	"fmt"
	"os"

	"github.com/arwinneil/country-cli/countries"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "country-cli",
	Short: "list countries and information about them",

	Run: func(cmd *cobra.Command, args []string) {

		europe := countries.FetchByRegion("Europe")
		northAmerica := countries.FetchBySubRegion("North America")

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
