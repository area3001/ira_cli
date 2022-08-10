package cmd

import (
	"fmt"
	"github.com/area3001/goira/sdk"
	"github.com/logrusorgru/aurora/v3"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure the IRA devices",
}

var configSetCmd = &cobra.Command{
	Use:   "set <selector> <parameter> <value>",
	Short: "Set a parameter for a device",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		i, err := strconv.Atoi(args[2])
		if err != nil {
			log.Panicln(err)
		}

		devs, err := client.Devices.Select(args[0])
		if err != nil {
			log.Panicln(aurora.Red(err))
		}

		devs.Perform(func(dev *sdk.Device) {
			if err := dev.SetConfig(args[1], byte(i)); err != nil {
				fmt.Println(dev.Meta.MAC, ":\t", aurora.Red("ERR\t"), aurora.Red(err.Error()))
				return
			}

			fmt.Println(dev.Meta.MAC, ":\t", aurora.Green("OK"))
		})
	},
}

var configNameCmd = &cobra.Command{
	Use:   "name <device> <device_name>",
	Short: "Assign the name for the device",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		dev, err := client.Devices.Device(args[0])
		if err != nil {
			log.Panicln(err)
		}

		if err := dev.SetName(args[1]); err != nil {
			log.Panicln(err)
		}
	},
}

func init() {
	configCmd.AddCommand(configSetCmd, configNameCmd)

	rootCmd.AddCommand(configCmd)
}
