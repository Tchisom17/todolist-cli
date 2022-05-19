package cmd

import (
	"github.com/tchisom17/todolist-cli/utils"
	"strconv"

	"github.com/spf13/cobra"
)

// undoneCmd represents the undone command
var undoneCmd = &cobra.Command{
	Use:   "undone",
	Short: "Mark a task as undone",
	Long: `Marks a task as undone`,
	Run: func(cmd *cobra.Command, args []string) {
		index, err := strconv.Atoi(args[0])
		if err != nil {
			panic(any(err))
		}
		utils.Todo.Undone(index)
	},
}

func init() {
	rootCmd.AddCommand(undoneCmd)
}
