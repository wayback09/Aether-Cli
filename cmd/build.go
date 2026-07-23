package cmd

import (
	"fmt"
	"os"

	"github.com/wayback09/aether-cli/pkg/manifest"
	"github.com/wayback09/aether-cli/pkg/packager"
)

func RunBuild(args []string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("error getting current directory: %w", err)
	}

	fmt.Println("Validating extension...")
	m, err := manifest.Validate(cwd)
	if err != nil {
		return fmt.Errorf("❌ Validation Failed. Cannot build: %w", err)
	}

	fmt.Printf("Packaging %s v%s...\n", m.Name, m.Version)
	outPath, err := packager.Build(cwd, m)
	if err != nil {
		return fmt.Errorf("❌ Build Failed: %w", err)
	}

	fmt.Printf("✅ Build Successful: %s\n", outPath)
	return nil
}
