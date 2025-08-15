/*
Copyright © 2025 Daniel Gimenez danielhgimenez027@gmail.com
*/
package complete

import (
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/DanielHGimenez/yell/src/alert"
	"github.com/spf13/cobra"
)

// commandCmd represents the command command
var commandCmd = &cobra.Command{
	Use:     "command [flags] -- args...",
	Short:   "Yell when the command execution is completed",
	Long:    `Notify when the command execution is completed.`,
	Args:    cobra.MinimumNArgs(1),
	Example: "yell on complete command -- sleep 10",
	Run: func(cmd *cobra.Command, args []string) {
		commandExecution := exec.Command(args[0], args[1:]...)

		output, err := cmd.Flags().GetBool("output")
		if err != nil {
			log.Fatal("could not read the 'output' flag: ", err)
		}
		if output {
			commandExecution.Stdout = os.Stdout
			commandExecution.Stderr = os.Stderr
		}

		timeout, err := cmd.Flags().GetInt16("timeout")
		if err != nil {
			log.Fatal("could not read the 'timeout' flag: ", err)
		}

		if err := commandExecution.Start(); err != nil {
			log.Fatal("could not run command: ", err)
		}

		if timeout > 0 {
			finished := make(chan string, 1)

			go func() {
				commandExecution.Wait()
				finished <- "done"
			}()

			select {
			case <-finished:
				// do nothing
			case <-time.After(time.Duration(timeout) * time.Millisecond):
				commandExecution.Process.Kill()
			}
		} else {
			commandExecution.Wait()
		}
		alert.Execute(cmd.Flags())
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// commandCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// commandCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	commandCmd.Flags().Int16P("timeout", "t", 0, "Timeout (in milliseconds) to wait for command completion. If the time is exceeded, the command is terminated and an alert is triggered.")
	commandCmd.Flags().BoolP("output", "o", false, "Print the output of the command.")
}
