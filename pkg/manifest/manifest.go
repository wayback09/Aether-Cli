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
