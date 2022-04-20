/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/area3001/goira/comm"
	"github.com/area3001/goira/sdk"
	"github.com/nats-io/nats.go"
	"log"

	"github.com/spf13/cobra"
)

// keysCmd represents the keys command
var keysCmd = &cobra.Command{
	Use:   "keys",
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

		keys, err := client.Devices.Keys()
		if err != nil {
			log.Panicln(err)
		}

		for _, v := range keys {
			fmt.Println(v)
		}
	},
}

func init() {
	devicesCmd.AddCommand(keysCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// keysCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// keysCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
