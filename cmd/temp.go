package cmd

import (
	"fmt"

	netatmo2 "github.com/mariusbreivik/netatmo/api/netatmo"
	"github.com/mariusbreivik/netatmo/internal/netatmo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ttacon/chalk"
)

var indoor bool
var outdoor bool

// tempCmd is the command for retrieving temperature
var tempCmd = &cobra.Command{
	Use:     "temp",
	Short:   "read temperature data from netatmo station",
	Long:    `read temperature data from netatmo station`,
	Example: "netatmo temp indoor",
	RunE: func(cmd *cobra.Command, args []string) error {
		netatmoClient, err := netatmo.NewClient(netatmo.Config{
			ClientID:     viper.GetString("netatmo.clientID"),
			ClientSecret: viper.GetString("netatmo.clientSecret"),
			Username:     viper.GetString("netatmo.username"),
			Password:     viper.GetString("netatmo.password"),
		})

		if err != nil {
			return err
		}

		if indoor {
			printIndoorTemp(netatmoClient.GetStationData())
		} else if outdoor {
			printOutdoorTemp(netatmoClient.GetStationData())
		} else {
			fmt.Println(cmd.UsageString())
		}

		return nil
	},
}

func printOutdoorTemp(stationData netatmo2.StationData) {
	fmt.Println("Station name: ", stationData.Body.Devices[0].StationName)
	fmt.Println("Temperature outdoor:", chalk.Green, stationData.Body.Devices[0].Modules[0].DashboardData.Temperature, chalk.Reset)

}

func printIndoorTemp(stationData netatmo2.StationData) {
	fmt.Println("Station name: ", stationData.Body.Devices[0].StationName)
	fmt.Println("Temperature indoor:", chalk.Red, stationData.Body.Devices[0].DashboardData.Temperature, chalk.Reset)
}

func init() {
	rootCmd.AddCommand(tempCmd)

	tempCmd.Flags().BoolVarP(&indoor, "indoor", "i", false, "netatmo temp -i|--indoor")
	tempCmd.Flags().BoolVarP(&outdoor, "outdoor", "o", false, "netatmo temp -o|--outdoor")

}
