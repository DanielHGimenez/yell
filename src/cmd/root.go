/*
Copyright © 2025 Daniel Gimenez danielhgimenez027@gmail.com
*/
package cmd

import (
	"os"

	"github.com/DanielHGimenez/yell/src/alert"
	"github.com/DanielHGimenez/yell/src/cmd/config"
	"github.com/DanielHGimenez/yell/src/cmd/on"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "yell",
	Short: "A smart command-line tool that alerts you when a task completes, so you can focus on what matters.",
	Long: `A smart command-line tool that alerts you when a task completes, so you can focus on what matters.

Never waste time manually watching tasks again! Yell monitors:
✅ Commands – Get notified when a long-running command finishes.
✅ Processes – Know when a background process exits.
✅ Dynamic Output – Detect changes in command output (great for waiting on conditions).

Just set it and forget it—get a notification when it’s done. 🚀`,

	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	Run: func(cmd *cobra.Command, args []string) {
		alert.Execute(cmd.Flags())
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	RootCmd.AddCommand(on.OnCmd)
	RootCmd.AddCommand(config.ConfigCmd)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.yell.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	alert.ConfigureFlags(RootCmd.Flags())
}
