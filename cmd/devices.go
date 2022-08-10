package cmd

import (
	"fmt"
	"github.com/area3001/goira/devices"
	"github.com/area3001/goira/sdk"
	"github.com/logrusorgru/aurora/v3"
	"github.com/mergestat/timediff"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strings"
)

var devicesCmd = &cobra.Command{
	Use:   "devices",
	Short: "IRA device management",
	Long:  `Manage ira devices`,
}

var forgetCmd = &cobra.Command{
	Use:   "forget <selector>",
	Short: "Forget the device",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		devs, err := client.Devices.Select(args[0])
		if err != nil {
			log.Panicln(aurora.Red(err))
		}

		devs.Perform(func(dev *sdk.Device) {
			if err := client.Devices.Forget(dev.Meta.MAC); err != nil {
				fmt.Println(dev.Meta.MAC, ":\t", aurora.Red("ERR\t"), aurora.Red(err.Error()))
				return
			}

			fmt.Println(dev.Meta.MAC, ":\t", aurora.Green("OK"))
		})

		return client.Devices.Forget(args[0])
	},
}

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "!!! Clean Everything !!!",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		keys, err := client.Devices.Keys()
		if err != nil {
			log.Panicln(aurora.Red(err))
		}

		for _, key := range keys {
			if err := client.Devices.Forget(key); err != nil {
				fmt.Println(key, ":\t", aurora.Red("ERR\t"), aurora.Red(err.Error()))
				continue
			}

			fmt.Println(key, ":\t", aurora.Green("OK"))
		}

		return nil
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
		table.SetAutoWrapText(false)
		table.SetHeader([]string{"Mac", "Name", "Mode", "Version", "Last Seen", "Config"})

		for _, v := range devs {
			if v.Meta == nil {
				continue
			}

			config := []string{
				fmt.Sprintf("pixel_length=%d", v.Meta.Config[devices.PixelLengthVar]),
				fmt.Sprintf("fx=%d", v.Meta.Config[devices.FxVar]),
				fmt.Sprintf("fx_speed=%d", v.Meta.Config[devices.FxSpeedVar]),
				fmt.Sprintf("fx_xfade=%d", v.Meta.Config[devices.FxXfadeVar]),
				fmt.Sprintf("fx_fg=(%d, %d, %d)", v.Meta.Config[devices.FxForegroundRedVar], v.Meta.Config[devices.FxForegroundGreenVar], v.Meta.Config[devices.FxForegroundBlueVar]),
				fmt.Sprintf("fx_bg=(%d, %d, %d)", v.Meta.Config[devices.FxBackgroundRedVar], v.Meta.Config[devices.FxBackgroundGreenVar], v.Meta.Config[devices.FxBackgroundBlueVar]),
			}

			table.Append([]string{v.Meta.MAC, v.Meta.Name, fmt.Sprintf("%d", v.Meta.Mode), fmt.Sprintf("%d", v.Meta.Version), timediff.TimeDiff(v.Meta.LastBeat), strings.Join(config, ", ")})
		}
		table.SetFooter([]string{"", "", "", "", "total", fmt.Sprintf("%d", len(devs))})
		table.Render()
	},
}

func init() {
	devicesCmd.AddCommand(deviceListCmd, syncCmd, forgetCmd, resetCmd, cleanCmd)

	rootCmd.AddCommand(devicesCmd)
}
