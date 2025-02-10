package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Younes-khadraoui/azrb/config"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Show current configuration",
	Long:  "Display the current storage directory and default editor settings, with instructions on how to modify them.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Current Configuration:")
		fmt.Println(" Storage Directory:", config.StorageDir)
		fmt.Println(" Default Editor:", config.Editor)
		fmt.Println("")
		home, err := os.UserHomeDir()
		if err != nil {
			home = "~"
		}
		fmt.Println("To change these settings, create a configuration file at", filepath.Join(home, ".azrb.yaml"))
		fmt.Println("Example contents:")
		fmt.Println(`storageDir: "/path/to/your/cheatsheets"
editor: "your_preferred_editor"`)
		fmt.Println("")
		fmt.Println("Alternatively, set the environment variables AZRB_STORAGE_DIR and AZRB_EDITOR.")
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
