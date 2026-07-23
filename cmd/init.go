package cmd

import (
	"fmt"
	"os"

	"github.com/wayback09/aether-cli/pkg/scaffold"
)

func RunInit(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("usage: aether init <extension-name> <extension-id>")
	}

	name := args[0]
	id := args[1]

	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("error getting current directory: %w", err)
	}

	targetDir := cwd + "/" + name
	if err := os.Mkdir(targetDir, 0755); err != nil {
		if os.IsExist(err) {
			return fmt.Errorf("directory %s already exists", name)
		}
		return fmt.Errorf("failed to create directory: %w", err)
	}

	fmt.Printf("Initializing extension '%s' (%s)...\n", name, id)
	if err := scaffold.CreateExtension(targetDir, name, id); err != nil {
		return fmt.Errorf("❌ Failed to initialize extension: %w", err)
	}

	fmt.Println("✅ Initialization complete!")
	fmt.Printf("cd %s && aether build\n", name)
	return nil
}
