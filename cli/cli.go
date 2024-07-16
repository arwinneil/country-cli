package cli

import (
	"fmt"
	"os"

	"github.com/arwinneil/country-cli/countries"
	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"
)

var pretty bool

var rootCmd = &cobra.Command{
	Use:   "country-cli",
	Short: "list countries and information about them",

	Run: func(cmd *cobra.Command, args []string) {

		europe, err := countries.FetchByRegion("Europe")

		if err != nil {
			fmt.Println("Failed to fetch Europe countries : %s", err.Error())
		}

		northAmerica, err := countries.FetchBySubRegion("North America")

		if err != nil {
			fmt.Println("Failed to fetch North America countries : %s", err.Error())
		}

		if pretty {
			t := table.NewWriter()
			t.SetOutputMirror(os.Stdout)
			t.AppendHeader(table.Row{"Name", "cca2", "cca3", "Region", "Subregion", "Currency", "Currency Symbol"})

			for _, c := range europe {
				t.AppendRow([]interface{}{c.Name, c.Cca2, c.Cca3, c.Region, c.Subregion, c.Currency, c.CurrencySymbo})
			}

			for _, c := range northAmerica {
				t.AppendRow([]interface{}{c.Name, c.Cca2, c.Cca3, c.Region, c.Subregion, c.Currency, c.CurrencySymbo})
			}
			t.Render()

			return
		}

		finalArray := append(europe, northAmerica...)

		fmt.Printf("%s\n", finalArray)

	},
}

func Execute() {
	rootCmd.Flags().BoolVar(&pretty, "pretty-print", false, "print table")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
