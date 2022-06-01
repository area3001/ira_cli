/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bufio"
	"github.com/area3001/goira/comm"
	"github.com/area3001/goira/sdk"
	"github.com/nats-io/nats.go"
	"github.com/spf13/cobra"
	"io"
	"log"
	"os"
)

// devicesCmd represents the devices command
var devicesCmd = &cobra.Command{
	Use:   "devices",
	Short: "IRA device management",
	Long:  `Manage ira devices`,
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure the IRA device",
}

var configSetCmd = &cobra.Command{
	Use:   "set device parameter value",
	Short: "read raw bytes sent as packets from stdin and send them to rgb",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := sdk.NewClient(&comm.NatsClientOpts{
			Root:             "area3001",
			NatsUrl:          server,
			NatsOptions:      []nats.Option{},
			JetStreamOptions: []nats.JSOpt{},
		})

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
			if err := dev.SendRgb(packet); err != nil {
				log.Println(err)
			}
		}
	},
}

func init() {
	devicesCmd.AddCommand(configCmd)
	rootCmd.AddCommand(devicesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// devicesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// devicesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
