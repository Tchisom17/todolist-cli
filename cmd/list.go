package cmd

import (
	"github.com/tchisom17/todolist-cli/utils"
	"log"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Display a list of all tasks",
	Long: `Display a list of all tasks`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("list called")
		utils.Todo.List()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

