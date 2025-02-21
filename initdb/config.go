package initdb

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

const configFile = "config.json"

type Config struct {
	DBPath string `json:"db_path"`
}

func LoadConfig() (Config, error) {
	var config Config
	file, err := os.ReadFile(configFile)
	if err != nil {
		return config, err
	}
	err = json.Unmarshal(file, &config)
	return config, err
}

func SaveConfig(dbPath string) error {
	// save path to conifg.json, also ensuring the db name is correctly set
	config := Config{DBPath: dbPath}
	data, err := json.MarshalIndent(config, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(configFile, data, 0644)
}

func CleanPath(path string) (string, error) {
	// handle empty path
	if path == "" {
		return "", fmt.Errorf("path cannot be empty")
	}

	// expand user-specific paths indicated by ~
	expandedPath, err := expandUser(path)
	if err != nil {
		return "", fmt.Errorf("failed to expand user path: %w", err)
	}

	// remove redundant separators and resolving .. elements
	cleanedPath := filepath.Clean(expandedPath)
	cleanedPath = strings.TrimSpace(cleanedPath)

	// ensure path is absolute
	absolutePath, err := filepath.Abs(cleanedPath)
	if err != nil {
		return "", fmt.Errorf("failed to get absolute path: %w", err)
	}
	// Add trailing slash if it's missing
	//    - Check if the path exists. If it doesn't exist, we assume it *will*
	//      be a directory.
	//    - If it exists, check if it *is* a directory.
	if !strings.HasSuffix(absolutePath, string(os.PathSeparator)) {
		fileInfo, err := os.Stat(absolutePath)
		if err == nil { // Path exists
			if fileInfo.IsDir() {
				absolutePath += string(os.PathSeparator)
			}
		} else if os.IsNotExist(err) {
			// Path does not exist, assume it will be a directory.
			absolutePath += string(os.PathSeparator)
		} else {
			// Some other error occurred during stat.
			return "", fmt.Errorf("failed to stat path: %w", err)
		}
	}
	// Convert separators to the OS-specific separator (between / and \)
	osSpecificPath := filepath.ToSlash(absolutePath) // Use forward slashes first

	if runtime.GOOS == "windows" {
		osSpecificPath = filepath.FromSlash(osSpecificPath) // Convert to backslashes on Windows
	}

	return osSpecificPath, nil
}

func expandUser(path string) (string, error) {
	// exits early if there is no ~ prefix
	if !strings.HasPrefix(path, "~") {
		return path, nil
	}

	// fetches full path of user's home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	// joins full path of home directory with path provided
	return filepath.Join(homeDir, path[1:]), nil
}
