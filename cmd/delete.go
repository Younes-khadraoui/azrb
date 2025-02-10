package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [cheatsheet name]",
	Short: "Delete an existing cheat sheet",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error fetching home directory:", err)
			return
		}
		storageDir := filepath.Join(homeDir, ".azrb")
		filePath := filepath.Join(storageDir, name+".md")

		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			fmt.Printf("Cheat sheet '%s' does not exist.\n", name)
			return
		}

		if err := os.Remove(filePath); err != nil {
			fmt.Println("Error deleting file:", err)
			return
		}
		fmt.Printf("Deleted cheat sheet: %s\n", filePath)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
