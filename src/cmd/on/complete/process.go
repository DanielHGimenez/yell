/*
Copyright © 2025 Daniel Gimenez danielhgimenez027@gmail.com
*/
package complete

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/DanielHGimenez/yell/src/alert"
	util "github.com/DanielHGimenez/yell/src/os"
	ps "github.com/mitchellh/go-ps"
	"github.com/spf13/cobra"
)

// processCmd represents the process command
var processCmd = &cobra.Command{
	Use:   "process",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var process *os.Process
		if cmd.Flag("pid").Changed {
			pid, err := cmd.Flags().GetString("pid")
			if err != nil {
				log.Fatal("could not read the 'pid' flag: ", err)
			}
			if pid == "" {
				log.Fatal("could not read the 'pid' flag.")
			}
			ipid, err := strconv.Atoi(pid)
			if err != nil {
				log.Fatal("could not read the 'pid' flag: ", err)
			}
			process, err = os.FindProcess(ipid)
			if err != nil {
				log.Fatal("could not find the process: ", err)
			}
		} else if cmd.Flag("name").Changed {
			name, err := cmd.Flags().GetString("name")
			if err != nil {
				log.Fatal("could not read the 'name' flag: ", err)
			}
			if name == "" {
				log.Fatal("could not read the 'name' flag.")
			}
			processList, err := ps.Processes()
			if err != nil {
				log.Fatal("could not get all processes: ", err)
			}
			for _, psprocess := range processList {
				if psprocess.Executable() == name {
					process, err = os.FindProcess(psprocess.Pid())
					if err != nil {
						log.Fatal("could not find the process: ", err)
					}
					break
				}
			}
		} else if cmd.Flag("contains").Changed {
			name, err := cmd.Flags().GetString("contains")
			if err != nil {
				log.Fatal("could not read the 'contains' flag: ", err)
			}
			if name == "" {
				log.Fatal("could not read the 'contains' flag.")
			}
			processList, err := ps.Processes()
			if err != nil {
				log.Fatal("could not get all processes: ", err)
			}
			for _, psprocess := range processList {
				if strings.Contains(psprocess.Executable(), name) {
					process, err = os.FindProcess(psprocess.Pid())
					if err != nil {
						log.Fatal("could not find the process: ", err)
					}
					break
				}
			}
		} else {
			log.Fatal("required flags were not set.")
		}
		if process == nil {
			log.Fatal("could not find the process")
		}
		for exists, _ := util.ProcessExists(process); exists; exists, _ = util.ProcessExists(process) {
			time.Sleep(100 * time.Millisecond)
		}
		alert.Execute(cmd.Flags())
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// processCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// processCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	processCmd.Flags().StringP("pid", "p", "", "Will lookup for a process with the same PID.")
	processCmd.Flags().StringP("name", "n", "", "Will lookup for a process with the exact name.")
	processCmd.Flags().StringP("contains", "c", "", "Will lookup for a process with a name containing the value.")
	processCmd.MarkFlagsOneRequired("pid", "name", "contains")
}
