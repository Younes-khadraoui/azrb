package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create [cheatsheet name]",
	Short: "Create a new cheat sheet",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error fetching home directory:", err)
			return
		}
		storageDir := filepath.Join(homeDir, ".azrb")
		os.MkdirAll(storageDir, os.ModePerm)
		filePath := filepath.Join(storageDir, name+".md")

		if _, err := os.Stat(filePath); err == nil {
			fmt.Printf("Cheat sheet '%s' already exists.\n", name)
			return
		}

		file, err := os.Create(filePath)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		file.Close()
		fmt.Printf("Created cheat sheet: %s\n", filePath)

		editor := os.Getenv("EDITOR")
		if editor == "" {

			if os.Getenv("OS") == "Windows_NT" {
				editor = "notepad"
			} else {
				editor = "nvim"
			}
		}
		cmdEditor := exec.Command(editor, filePath)
		cmdEditor.Stdin = os.Stdin
		cmdEditor.Stdout = os.Stdout
		cmdEditor.Stderr = os.Stderr
		if err := cmdEditor.Run(); err != nil {
			fmt.Println("Error opening editor:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
