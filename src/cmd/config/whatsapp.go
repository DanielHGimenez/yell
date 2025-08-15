package config

import (
	"log"
	"time"

	"github.com/DanielHGimenez/yell/src/integration/whatsapp"
	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var whatsappCmd = &cobra.Command{
	Use:   "whatsapp",
	Short: "Connect to the whatsapp account.",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := whatsapp.CreateClient()
		if err != nil {
			log.Fatal("could not login to whatsapp: ", err)
		}
		defer client.Disconnect()
		err = whatsapp.Login(client)
		if err != nil {
			log.Fatal("could not login to whatsapp: ", err)
		}
		time.Sleep(1 * time.Second) // Wait for the client to connect
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
