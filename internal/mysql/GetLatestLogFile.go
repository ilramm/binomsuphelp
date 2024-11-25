package internal

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// Filter log files based on pattern
func filterLogFiles(output string) []string {
	var filtered []string
	lines := strings.Split(output, "\n")
	// Define regex to match mysql_error_log.x where x is a number or none
	logPattern := regexp.MustCompile(`^mysql_error_log(\.(\d+))?$`)

	for _, line := range lines {
		line = strings.TrimSpace(line) // Remove leading/trailing spaces
		if logPattern.MatchString(line) {
			filtered = append(filtered, line)
		}
	}
	return filtered
}

// Extract the numeric suffix (if any) from the file name
func getLogFileSuffix(fileName string) int {
	// If no suffix (e.g., mysql_error_log), treat it as 0
	if !strings.Contains(fileName, ".") {
		return 0
	}

	// Split by "." and extract the last part
	parts := strings.Split(fileName, ".")
	if len(parts) > 1 {
		if num, err := strconv.Atoi(parts[len(parts)-1]); err == nil {
			return num
		}
	}

	return 0
}

// GetLatestLogFile returns the latest log file from a directory listing string
func GetLatestLogFile(dirListing string) (string, error) {
	// Filter relevant log files
	files := filterLogFiles(dirListing)
	if len(files) == 0 {
		return "", fmt.Errorf("no mysql_error_log files found")
	}

	// Sort log files by numeric suffix in descending order
	sort.Slice(files, func(i, j int) bool {
		iSuffix := getLogFileSuffix(files[i])
		jSuffix := getLogFileSuffix(files[j])
		return iSuffix > jSuffix
	})

	// Return the latest log file
	return files[0], nil
}
