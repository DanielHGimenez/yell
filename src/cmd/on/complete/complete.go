/*
Copyright © 2025 Daniel Gimenez danielhgimenez027@gmail.com
*/
package complete

import (
	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var CompleteCmd = &cobra.Command{
	Use:   "complete",
	Short: "Yell when the event is finished.",
	Long:  `Notify when the specified event is finished/completed.`,
}

func init() {
	CompleteCmd.AddCommand(commandCmd)
	CompleteCmd.AddCommand(processCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
