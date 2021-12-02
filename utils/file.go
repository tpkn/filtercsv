package utils

import (
	"errors"
	"os"
	"strings"
)

// FileHashify converts file lines into a map
func FileHashify(file_path string) (map[string]int, error) {
	content, err := os.ReadFile(file_path)
	if err != nil {
		return nil, err
	}
	
	var keys = strings.Split(string(content), "\n")
	var keys_map = make(map[string]int, len(keys))
	
	for _, k := range keys {
		key := strings.TrimSpace(k)
		if key != "" {
			keys_map[key] = 1
		}
	}
	
	// Throw error if map is empty
	if len(keys_map) == 0 {
		return nil, errors.New("seems like '" + file_path + "' file is empty")
	}
	
	return keys_map, nil
}

// FileExists checks if the file exists
func FileExists(file_path string) bool {
	if file_path == "" {
		return false
	}
	
	i, err := os.Stat(file_path)
	if os.IsNotExist(err) || i.IsDir() {
		return false
	}
	return true
}
