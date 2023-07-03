package utils

import (
"strings"
"path/filepath"

)

func GetFileType(filename string) string {
	// Get the file extension
	ext := filepath.Ext(filename)
	// Remove the dot from the extension
	return strings.ToLower(strings.TrimPrefix(ext, "."))
}