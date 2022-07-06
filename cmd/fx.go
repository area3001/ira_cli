package cmd

import (
	"fmt"
	"github.com/area3001/goira/core"
	"github.com/area3001/goira/sdk"
	"github.com/logrusorgru/aurora/v3"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strconv"
	"strings"
)

var fxCmd = &cobra.Command{
	Use:   "fx",
	Short: "Device Effects control",
}

var fxEnableCmd = &cobra.Command{
	Use:   "enable <selector>",
	Short: "Enable RGB mode for the selected devices",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		devs, err := client.Devices.Select(args[0])
		if err != nil {
			log.Panicln(aurora.Red(err))
		}

		devs.Perform(func(dev *sdk.Device) {
			if err := dev.SetMode(core.Modes[8]); err != nil {
				fmt.Println(dev.Meta.MAC, ":\t", aurora.Red("ERR\t"), aurora.Red(err.Error()))
				return
			}

			fmt.Println(dev.Meta.MAC, ":\t", aurora.Green("OK"))
		})

		_ = client.Devices.Sync()
	},
}

var listFxCmd = &cobra.Command{
	Use:   "list",
	Short: "list the available effects with their code",
	Run: func(cmd *cobra.Command, args []string) {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Code", "Name", "Description", "Parameters"})

		for _, v := range core.Effects {
			params := make([]string, len(v.AllowedParams))
			for idx, param := range v.AllowedParams {
				params[idx] = param.Name
			}

			table.Append([]string{fmt.Sprintf("%d", v.Code), v.Name, v.Description, strings.Join(params, ", ")})
		}

		table.Render()
	},
}

var setFxCmd = &cobra.Command{
	Use:   "set <selector> <effect_code> [<param>=<value>]*",
	Short: "set the specified effect",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		effectIdx, err := strconv.Atoi(args[1])
		if err != nil {
			log.Panicln(aurora.Red(err.Error()))
		}

		params := map[string]string{}
		for _, opt := range args[2:] {
			parts := strings.Split(opt, "=")
			if len(parts) != 2 {
				log.Panicln("invalid parameter", opt)
			}

			params[parts[0]] = parts[1]
		}

		fx, err := core.NewEffect(core.Effects[effectIdx], params)
		if err != nil {
			log.Panicln(aurora.Red(err))
		}

		devs, err := client.Devices.Select(args[0])
		if err != nil {
			log.Panicln(aurora.Red(err))
		}

		devs.Perform(func(dev *sdk.Device) {
			if err := dev.SendFx(fx); err != nil {
				fmt.Println(dev.Meta.MAC, ":\t", aurora.Red("ERR\t"), aurora.Red(err.Error()))
				return
			}

			fmt.Println(dev.Meta.MAC, ":\t", aurora.Green("OK"))
		})
	},
}

func init() {
	fxCmd.AddCommand(fxEnableCmd, listFxCmd, setFxCmd)
	rootCmd.AddCommand(fxCmd)
}
