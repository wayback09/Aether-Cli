package packager

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/wayback09/aether-cli/pkg/manifest"
)

// Build packages the extension directory into an .aex zip archive.
func Build(dirPath string, m *manifest.Manifest) (string, error) {
	outName := fmt.Sprintf("%s-%s.aex", m.ID, m.Version)
	outPath := filepath.Join(dirPath, outName)

	outFile, err := os.Create(outPath)
	if err != nil {
		return "", fmt.Errorf("failed to create output file: %w", err)
	}
	defer outFile.Close()

	zipWriter := zip.NewWriter(outFile)
	defer zipWriter.Close()

	err = filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		
		// Skip the output file itself
		if path == outPath {
			return nil
		}

		relPath, err := filepath.Rel(dirPath, path)
		if err != nil {
			return err
		}

		// Ignore .git, node_modules, and typical dev files
		if strings.HasPrefix(relPath, ".git") || 
			strings.HasPrefix(relPath, "node_modules") || 
			strings.HasSuffix(relPath, ".aex") ||
			strings.HasSuffix(relPath, ".zip") {
			return nil
		}

		if info.IsDir() {
			return nil // ZIP entries for directories are implicitly created if files exist, or we can add them manually.
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		// Create a file header inside the zip
		// Use forward slashes for ZIP paths
		zipPath := filepath.ToSlash(relPath)
		writer, err := zipWriter.Create(zipPath)
		if err != nil {
			return err
		}

		if _, err := io.Copy(writer, file); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return "", fmt.Errorf("failed to package extension: %w", err)
	}

	return outPath, nil
}
