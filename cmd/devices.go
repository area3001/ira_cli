package cmd

import (
	"fmt"
	"github.com/area3001/goira/sdk"
	"github.com/logrusorgru/aurora/v3"
	"github.com/mergestat/timediff"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"log"
	"os"
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

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Ask all devices to sync",
	Run: func(cmd *cobra.Command, args []string) {
		if err := client.Devices.Sync(); err != nil {
			log.Panicln(err)
		}
	},
}

var resetCmd = &cobra.Command{
	Use:   "reset <selector>",
	Short: "Reset the selected devices",
	Run: func(cmd *cobra.Command, args []string) {
		devs, err := client.Devices.Select(args[0])
		if err != nil {
			log.Panicln(aurora.Red(err))
		}

		devs.Perform(func(dev *sdk.Device) {
			if err := dev.Reset(500); err != nil {
				fmt.Println(dev.Meta.MAC, ":\t", aurora.Red("ERR\t"), aurora.Red(err.Error()))
				return
			}

			fmt.Println(dev.Meta.MAC, ":\t", aurora.Green("OK"))
		})
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
		table.SetHeader([]string{"Mac", "Name", "Mode", "Last Seen"})

		for _, v := range devs {
			if v.Meta == nil {
				continue
			}

			table.Append([]string{v.Meta.MAC, v.Meta.Name, v.Meta.Mode, timediff.TimeDiff(v.Meta.LastBeat)})
		}
		table.Render()
	},
}

func init() {
	devicesCmd.AddCommand(deviceListCmd, syncCmd, forgetCmd, resetCmd)

	rootCmd.AddCommand(devicesCmd)
}
