package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/wayback09/aether-cli/pkg/manifest"
)

func RunValidate(args []string) error {
	isTheme := false
	for _, arg := range args {
		if arg == "--theme" || arg == "-t" {
			isTheme = true
		}
	}

	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("error getting current directory: %w", err)
	}

	// Auto-detect: if manifest.json is absent but package.json is present, treat as theme
	if !isTheme {
		manifestPath := filepath.Join(cwd, "manifest.json")
		if _, err := os.Stat(manifestPath); os.IsNotExist(err) {
			pkgPath := filepath.Join(cwd, "package.json")
			if _, err := os.Stat(pkgPath); err == nil {
				isTheme = true
			}
		}
	}

	if isTheme {
		fmt.Println("Validating theme in:", cwd)
		t, err := manifest.ValidateTheme(cwd)
		if err != nil {
			return fmt.Errorf("❌ Theme Validation Failed: %w", err)
		}
		fmt.Printf("✅ Theme Validation Passed: %s (%s) v%s\n", t.Name, t.ID, t.Version)
	} else {
		fmt.Println("Validating extension in:", cwd)
		m, err := manifest.Validate(cwd)
		if err != nil {
			return fmt.Errorf("❌ Extension Validation Failed: %w", err)
		}
		fmt.Printf("✅ Extension Validation Passed: %s (%s) v%s\n", m.Name, m.ID, m.Version)
	}

	return nil
}
