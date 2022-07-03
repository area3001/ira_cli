package cmd

import (
	"fmt"
	"github.com/area3001/goira/core"
	"github.com/area3001/goira/sdk"
	"github.com/logrusorgru/aurora/v3"
	"github.com/mergestat/timediff"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strconv"
)

var devicesCmd = &cobra.Command{
	Use:   "devices",
	Short: "IRA device management",
	Long:  `Manage ira devices`,
}

var forgetCmd = &cobra.Command{
	Use:   "forget <device>",
	Short: "Forget the device",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return client.Devices.Forget(args[0])
	},
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure the IRA device",
}

var confModeCmd = &cobra.Command{
	Use:   "mode <device> <target_mode>",
	Short: "Set the device mode",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		targetMode, err := strconv.Atoi(args[1])
		if err != nil {
			log.Fatalln("invalid target mode format")
		}

		if targetMode < 0 || targetMode >= len(core.Modes) {
			log.Fatalln("target mode out of bounds")
		}

		mode := core.Modes[targetMode]

		devs, err := client.Devices.Select(args[0])
		if err != nil {
			log.Panicln(aurora.Red(err))
		}

		devs.Perform(func(dev *sdk.Device) {
			fmt.Print(dev.Meta.MAC, " ... ")

			if err := dev.SetMode(mode); err != nil {
				fmt.Println(aurora.Red(err.Error()))
				return
			}

			fmt.Println(aurora.Green("OK"))
		})

		_ = client.Devices.Sync()
		return nil
	},
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

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Ask all devices to sync",
	Run: func(cmd *cobra.Command, args []string) {
		if err := client.Devices.Sync(); err != nil {
			log.Panicln(err)
		}
	},
}

var deviceShowCmd = &cobra.Command{
	Use:   "show <device>",
	Short: "show the device details",
	Run: func(cmd *cobra.Command, args []string) {
		devs, err := client.Devices.List()
		if err != nil {
			log.Panicln(err)
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Mac", "Mode", "Last Seen"})

		for _, v := range devs {
			if v.Meta == nil {
				continue
			}

			table.Append([]string{v.Meta.MAC, v.Meta.Mode, timediff.TimeDiff(v.Meta.LastBeat)})
		}
		table.Render()
	},
}

var deviceListCmd = &cobra.Command{
	Use:   "list",
	Short: "list the devices",
	Long:  `List the devices showing their MAC, Mode and when they were last seen`,
	Run: func(cmd *cobra.Command, args []string) {
		devs, err := client.Devices.List()
		if err != nil {
			log.Panicln(err)
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Mac", "Mode", "Last Seen"})

		for _, v := range devs {
			if v.Meta == nil {
				continue
			}

			table.Append([]string{v.Meta.MAC, v.Meta.Mode, timediff.TimeDiff(v.Meta.LastBeat)})
		}
		table.Render()
	},
}

func init() {
	configCmd.AddCommand(configSetCmd, confModeCmd)

	devicesCmd.AddCommand(configCmd, deviceListCmd, syncCmd, forgetCmd)

	rootCmd.AddCommand(devicesCmd)
}
