/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/area3001/goira/comm"
	"github.com/area3001/goira/sdk"
	"github.com/nats-io/nats.go"
	"log"

	"github.com/spf13/cobra"
)

// blinkCmd represents the blink command
var blinkCmd = &cobra.Command{
	Use:   "blink [-n times] device",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := sdk.NewClient(&comm.NatsClientOpts{
			Root:             "area3001",
			NatsUrl:          server,
			NatsOptions:      []nats.Option{},
			JetStreamOptions: []nats.JSOpt{},
		})

		if err != nil {
			log.Panicln(err)
		}

		dev, err := client.Devices.Device(args[0])
		if err != nil {
			log.Panicln(err)
		}

		if dev == nil {
			log.Panicln("no device found with key " + args[0])
		}

		if err := dev.Blink(times); err != nil {
			log.Panicln(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(blinkCmd)
	blinkCmd.Flags().IntVarP(&times, "times", "n", 5, "the number of times the debug led should flash")
}
