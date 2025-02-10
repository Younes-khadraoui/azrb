package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"azrb/config"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "azrb [cheatsheet name]",
	Short: "azrb is a CLI cheat sheet manager",
	Long:  "azrb lets you create, view, edit, delete, list cheat sheets and view configuration.",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			cheatSheetName := args[0]
			filePath := filepath.Join(config.StorageDir, cheatSheetName+".md")
			data, err := os.ReadFile(filePath)
			if err != nil {
				fmt.Printf("Error reading cheat sheet '%s': %v\n", cheatSheetName, err)
				return
			}
			fmt.Println(string(data))
			return
		}

		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
