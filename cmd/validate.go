package cmd

import (
	"fmt"
	"os"

	"github.com/wayback09/aether-cli/pkg/manifest"
)

func RunValidate(args []string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("error getting current directory: %w", err)
	}

	fmt.Println("Validating extension in:", cwd)
	m, err := manifest.Validate(cwd)
	if err != nil {
		return fmt.Errorf("❌ Validation Failed: %w", err)
	}

	fmt.Printf("✅ Validation Passed: %s (%s) v%s\n", m.Name, m.ID, m.Version)
	return nil
}
