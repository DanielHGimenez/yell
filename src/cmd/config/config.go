package config

import (
	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Configuration options.",
}

func init() {
	ConfigCmd.AddCommand(whatsappCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
