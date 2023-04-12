package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
)

func extractPackage(packagePath string, outputPath string) error {
	// Set default output path if none provided
	if outputPath == "" {
		wd, err := os.Getwd()
		if err != nil {
			return err
		}
		outputPath = wd
	}

	// Create temporary directory to extract files to
	tmpDir, err := os.MkdirTemp("", "unitypkg")
	if err != nil {
		return errors.New("Failed to create temporary directory: " + err.Error())
	}
	defer os.RemoveAll(tmpDir)

	// Extract entire package to temporary directory
	cmd := exec.Command("tar", "xzf", packagePath, "-C", tmpDir)
	if err := cmd.Run(); err != nil {
		return errors.New("Failed to extract package to temp directory: " + err.Error())
	}

	// Extract each file to final destination
	files, err := os.ReadDir(tmpDir)
	if err != nil {
		return errors.New("Failed to read temp directory: " + err.Error())
	}
	for _, file := range files {
		assetEntryDir := filepath.Join(tmpDir, file.Name())
		pathnamePath := filepath.Join(assetEntryDir, "pathname")
		assetPath := filepath.Join(assetEntryDir, "asset")

		// Check if file has required info to extract
		if _, err := os.Stat(pathnamePath); os.IsNotExist(err) {
			fmt.Printf("WARNING: Skipping '%s' as it does not contain a pathname file.\n", file.Name())
			continue
		}

		// Get pathname from pathname file
		pathnameBytes, err := os.ReadFile(pathnamePath)
		if err != nil {
			return err
		}
		pathname := string(pathnameBytes) // Remove newline
		if os.PathSeparator == '\\' {     // Replace reserved characters on Windows
			re := regexp.MustCompile(`[<>:"|?*]`)
			pathname = re.ReplaceAllString(pathname, "_")
		}

		// Determine output path for asset
		assetOutPath := filepath.Join(outputPath, pathname)
		if !filepath.HasPrefix(filepath.Clean(assetOutPath), filepath.Clean(outputPath)) {
			fmt.Printf("WARNING: Skipping '%s' as '%s' is outside of '%s'.\n", file.Name(), assetOutPath, outputPath)
			continue
		}

		// Check if the asset file exists before moving
		if _, err := os.Stat(assetPath); os.IsNotExist(err) {
			fmt.Printf("WARNING: Skipping '%s' as asset file is missing.\n", file.Name())
			continue
		}

		// Move asset file to output path
		fmt.Printf("Extracting '%s' as '%s'\n", file.Name(), pathname)
		if err := os.MkdirAll(filepath.Dir(assetOutPath), os.ModePerm); err != nil {
			return errors.New("Failed to create directory for asset: " + err.Error())
		}
		if err := os.Rename(assetPath, assetOutPath); err != nil {
			return errors.New("Failed to move asset to output path: " + err.Error())
		}
	}

	return nil
}
