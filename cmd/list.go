package cmd

import (
	"fmt"
	"os"
	"strings"

	"azrb/config"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all cheat sheets",
	Run: func(cmd *cobra.Command, args []string) {
		dir := config.StorageDir
		files, err := os.ReadDir(dir)
		if err != nil {
			fmt.Println("Error reading storage directory:", err)
			return
		}

		var found bool
		fmt.Println("Cheat sheets:")
		for _, file := range files {
			if !file.IsDir() && strings.HasSuffix(file.Name(), ".md") {
				name := strings.TrimSuffix(file.Name(), ".md")
				fmt.Println(" -", name)
				found = true
			}
		}
		if !found {
			fmt.Println(" No cheat sheets found.")
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
