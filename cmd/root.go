package cmd

import (
	"github.com/spf13/cobra"
)

//var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "todolist-cli",
	Short: "Display your todo list",
	Long: `Todolist-cli is a CLI tool to display your todo list.
It is a simple todo list application that can be used to manage your todo list.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
