package manifest

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// Manifest represents the structure of manifest.json
type Manifest struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Version     string   `json:"version"`
	Author      string   `json:"author"`
	Description string   `json:"description"`
	Main        string   `json:"main"`
	API         string   `json:"api"`
	Permissions []string `json:"permissions"`
}

// Validate reads the manifest.json from the given directory and validates it.
func Validate(dirPath string) (*Manifest, error) {
	manifestPath := filepath.Join(dirPath, "manifest.json")
	
	fileBytes, err := os.ReadFile(manifestPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, errors.New("manifest.json not found in the current directory")
		}
		return nil, fmt.Errorf("error reading manifest.json: %w", err)
	}

	var m Manifest
	if err := json.Unmarshal(fileBytes, &m); err != nil {
		return nil, fmt.Errorf("invalid manifest.json syntax: %w", err)
	}

	// Validate required fields
	var missing []string
	if m.ID == "" {
		missing = append(missing, "id")
	}
	if m.Name == "" {
		missing = append(missing, "name")
	}
	if m.Version == "" {
		missing = append(missing, "version")
	}
	if m.Main == "" {
		missing = append(missing, "main")
	}
	if m.API == "" {
		missing = append(missing, "api")
	}

	if len(missing) > 0 {
		return &m, fmt.Errorf("manifest is missing required fields: %v", missing)
	}

	return &m, nil
}

// ThemeManifest represents the structure of package.json for themes
type ThemeManifest struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Version     string `json:"version"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Icon        string `json:"icon,omitempty"`
	CSS         string `json:"css,omitempty"`
	Overwrite   string `json:"overwrite,omitempty"`
}

// ValidateTheme reads the package.json from the given directory and validates it as a theme.
func ValidateTheme(dirPath string) (*ThemeManifest, error) {
	pkgPath := filepath.Join(dirPath, "package.json")
	
	fileBytes, err := os.ReadFile(pkgPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, errors.New("package.json not found in the current directory")
		}
		return nil, fmt.Errorf("error reading package.json: %w", err)
	}

	var t ThemeManifest
	if err := json.Unmarshal(fileBytes, &t); err != nil {
		return nil, fmt.Errorf("invalid package.json syntax: %w", err)
	}

	// Validate required fields
	var missing []string
	if t.ID == "" {
		missing = append(missing, "id")
	}
	if t.Name == "" {
		missing = append(missing, "name")
	}
	if t.Version == "" {
		missing = append(missing, "version")
	}

	if len(missing) > 0 {
		return &t, fmt.Errorf("theme package.json is missing required fields: %v", missing)
	}

	// Set defaults
	if t.CSS == "" {
		t.CSS = "theme.css"
	}
	if t.Overwrite == "" {
		t.Overwrite = "overwrite.json"
	}

	// Check if CSS file exists
	cssPath := filepath.Join(dirPath, t.CSS)
	if _, err := os.Stat(cssPath); os.IsNotExist(err) {
		return &t, fmt.Errorf("required stylesheet %q not found", t.CSS)
	}

	return &t, nil
}
