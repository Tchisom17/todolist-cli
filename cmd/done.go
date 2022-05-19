package cmd

import (
	"github.com/tchisom17/todolist-cli/utils"
	"strconv"

	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Mark a task as done",
	Long: `Mark a task as done`,
	Run: func(cmd *cobra.Command, args []string) {
		index, err := strconv.Atoi(args[0])
		if err != nil {
			panic(any(err))
		}
		utils.Todo.Done(index)
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
