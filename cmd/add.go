package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tchisom17/todolist-cli/utils"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add todo to list",
	Long: `Add todo to list.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			utils.Todo.Add(utils.AddTodoWithTerm())
		}else {
			utils.Todo.Add(args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.PersistentFlags().String("add", "", "Add todo to list.")
}
