package nginx

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func filterLogFiles2(output string) []string {
	var filtered []string
	lines := strings.Split(output, "\n")
	logPattern := regexp.MustCompile(`^binom.error.log(\.(\d+))?$`)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if logPattern.MatchString(line) {
			filtered = append(filtered, line)
		}
	}
	return filtered
}

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
	files := filterLogFiles2(dirListing)
	if len(files) == 0 {
		return "", fmt.Errorf("no binom.error.log files found")
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
