package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var blinkCmd = &cobra.Command{
	Use:   "blink <device>",
	Short: "Blink the debug led",
	Long:  `Blink the debug led any number of times. If no times parameter is provided, the led will blink 5 times.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
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
