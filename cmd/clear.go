package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tchisom17/todolist-cli/utils"
)

// clearCmd represents the clear command
var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Remove all todos set to done",
	Long: `Remove all todos set to done.`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.Todo.Clear()
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)
}

