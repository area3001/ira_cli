package cmd

import (
	"bufio"
	"fmt"
	"github.com/area3001/goira/core"
	"github.com/area3001/goira/sdk"
	"github.com/logrusorgru/aurora/v3"
	"github.com/spf13/cobra"
	"image/color"
	"io"
	"log"
	"os"
)

var rgbCmd = &cobra.Command{
	Use:   "rgb",
	Short: "Device RGB control",
}

var rgbRawCmd = &cobra.Command{
	Use:   "raw <device>",
	Short: "read raw bytes sent as packets from stdin and send them to rgb",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dev, err := client.Devices.Device(args[0])
		if err != nil {
			log.Panicln(err)
		}

		reader := bufio.NewReader(os.Stdin)
		finish := false
		for finish {
			packet, err := reader.ReadBytes('\n')
			if err != nil {
				if err == io.EOF {
					finish = true
				} else {
					log.Panic(err)
				}
			}

			if len(packet) <= 4 {
				log.Printf("Invalid packet length: %d\n", len(packet))
				continue
			}

			// -- send the packet
			if err := dev.SendRgbRaw(packet); err != nil {
				log.Println(err)
			}
		}
	},
}

var rgbSetCmd = &cobra.Command{
	Use:   "set <device> <hex_color, ...>",
	Short: "read raw bytes sent as packets from stdin and send them to rgb",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		// -- get the offset
		offset, _ := cmd.Flags().GetUint16("offset")

		devs, err := client.Devices.Select(args[0])
		if err != nil {
			log.Panicln(aurora.Red(err))
		}

		devs.Perform(func(dev *sdk.Device) {
			colors := make([]color.RGBA, len(args)-1)
			for idx, hexColor := range args[1:] {
				c, err := core.ParseHexColor(hexColor)
				if err != nil {
					msg := fmt.Sprintf("unable to parse color %d from %s to a valid color: %v\n", idx+1, hexColor, err)
					fmt.Println(dev.Meta.MAC, ":\t", aurora.Red("ERR\t"), aurora.Red(msg))
					continue
				}
				colors[idx] = c
			}

			if err := dev.SendRgbPixels(offset, colors); err != nil {
				fmt.Println(dev.Meta.MAC, ":\t", aurora.Red("ERR\t"), aurora.Red(err.Error()))
				return
			}

			fmt.Println(dev.Meta.MAC, ":\t", aurora.Green("OK"))
		})
	},
}

func init() {
	rgbCmd.AddCommand(rgbRawCmd)

	rgbSetCmd.LocalFlags().IntP("offset", "o", 0, "the pixel offset")

	rgbCmd.AddCommand(rgbSetCmd)

	rootCmd.AddCommand(rgbCmd)
}
