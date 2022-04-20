/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/area3001/goira/comm"
	"github.com/area3001/goira/sdk"
	"github.com/nats-io/nats.go"
	"github.com/spf13/cobra"
	"log"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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

		devs, err := client.Devices.List()
		if err != nil {
			log.Panicln(err)
		}

		for _, v := range devs {
			if v.Meta == nil {
				continue
			}

			fmt.Printf("%s %s %s\n", v.Meta.MAC, v.Meta.Mode, v.Meta.LastBeat)
		}
	},
}

func init() {
	devicesCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
