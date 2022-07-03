package cmd

import (
	"bufio"
	"github.com/spf13/cobra"
	"io"
	"log"
	"os"
)

var fxCmd = &cobra.Command{
	Use:   "fx",
	Short: "Device Effects control",
}

var listFxCmd = &cobra.Command{
	Use:   "list",
	Short: "list the available effects with their code",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var setFxCmd = &cobra.Command{
	Use:   "set <device>",
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

//var rgbSetCmd = &cobra.Command{
//	Use:   "set <device>",
//	Short: "read raw bytes sent as packets from stdin and send them to rgb",
//	Args:  cobra.MinimumNArgs(2),
//	Run: func(cmd *cobra.Command, args []string) {
//		dev, err := client.Devices.Device(args[0])
//		if err != nil {
//			log.Panicln(err)
//		}
//
//		// -- construct the packet
//		bytesPerPixel := len(args[0])
//		if bytesPerPixel != 6 && bytesPerPixel != 8 {
//			log.Fatalf("wrong number of pixel channels\n")
//		}
//
//		buf := new(bytes.Buffer)
//		binary.Write(buf, binary.LittleEndian, uint16(rgbOffset))
//		binary.Write(buf, binary.LittleEndian, uint16(len(args)))
//
//		for idx, arg := range args {
//			b, err := hex.DecodeString(arg)
//			if err != nil {
//				log.Fatalf("Data %s for pixel %d is invalid: %s\n", arg, idx, err)
//			}
//
//			if len(b) != bytesPerPixel {
//				log.Fatalf("Data %s for pixel %d is invalid: inconsistent number of bytes per pixel", arg, idx)
//			}
//
//			buf.Write(b)
//		}
//
//		// -- send the packet
//		if err := dev.SendRgbRaw(buf.Bytes()); err != nil {
//			log.Println(err)
//		}
//	},
//}

func init() {
	//rgbCmd.AddCommand(rgbRawCmd)
	//
	//rgbSetCmd.LocalFlags().IntVarP(&rgbOffset, "offset", "o", 0, "the pixel offset")
	//rgbCmd.AddCommand(rgbSetCmd)
	//
	//rootCmd.AddCommand(rgbCmd)
}
