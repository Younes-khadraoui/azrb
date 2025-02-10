package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/Younes-khadraoui/azrb/config"

	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit [cheatsheet name]",
	Short: "Edit an existing cheat sheet",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		filePath := filepath.Join(config.StorageDir, name+".md")
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			fmt.Printf("Cheat sheet '%s' does not exist. Use 'azrb create %s' to create it.\n", name, name)
			return
		}
		editorCmd := exec.Command(config.Editor, filePath)
		editorCmd.Stdin = os.Stdin
		editorCmd.Stdout = os.Stdout
		editorCmd.Stderr = os.Stderr
		if err := editorCmd.Run(); err != nil {
			fmt.Println("Error opening editor:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}
