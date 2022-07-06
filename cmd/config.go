package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure the IRA devices",
}

var configSetCmd = &cobra.Command{
	Use:   "set <device> <parameter> <value>",
	Short: "Set a parameter for a device",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		dev, err := client.Devices.Device(args[0])
		if err != nil {
			log.Panicln(err)
		}

		if err := dev.SetConfig(args[1], args[2]); err != nil {
			log.Panicln(err)
		}
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

		if err := dev.SetConfig(args[1], args[2]); err != nil {
			log.Panicln(err)
		}
	},
}

func init() {
	configCmd.AddCommand(configSetCmd, configNameCmd)

	rootCmd.AddCommand(configCmd)
}
