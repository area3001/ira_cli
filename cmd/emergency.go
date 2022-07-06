package cmd

import (
	"fmt"
	"github.com/area3001/goira/core"
	"github.com/area3001/goira/sdk"
	"github.com/logrusorgru/aurora/v3"
	"github.com/spf13/cobra"
	"log"
)

var emergencyCmd = &cobra.Command{
	Use:   "emergency",
	Short: "Enable emergency mode",
	Run: func(cmd *cobra.Command, args []string) {
		devs, err := client.Devices.Select("all")
		if err != nil {
			log.Panicln(aurora.Red(err))
		}

		devs.Perform(func(dev *sdk.Device) {
			if err := dev.SetMode(core.Modes[11]); err != nil {
				fmt.Println(dev.Meta.MAC, ":\t", aurora.Red("ERR\t"), aurora.Red(err.Error()))
				return
			}

			fmt.Println(dev.Meta.MAC, ":\t", aurora.Green("OK"))
		})

		_ = client.Devices.Sync()
	},
}

func init() {
	rootCmd.AddCommand(emergencyCmd)
}
