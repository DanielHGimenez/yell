/*
Copyright © 2025 Daniel Gimenez danielhgimenez027@gmail.com
*/
package on

import (
	"github.com/DanielHGimenez/yell/src/alert"
	"github.com/DanielHGimenez/yell/src/cmd/on/complete"
	"github.com/spf13/cobra"
)

// onCmd represents the on command
var OnCmd = &cobra.Command{
	Use:   "on",
	Short: "Definitions for what happens \"on\" some type of event.",
}

func init() {
	OnCmd.AddCommand(complete.CompleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// onCmd.PersistentFlags().String("foo", "", "A help for foo")
	alert.ConfigureFlags(OnCmd.PersistentFlags())

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// onCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
