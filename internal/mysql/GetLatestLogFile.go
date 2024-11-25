package mysql

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func filterLogFiles(output string) []string {
	var filtered []string
	lines := strings.Split(output, "\n")
	logPattern := regexp.MustCompile(`^mysql_error.log(\.(\d+))?$`)

	for _, line := range lines {
		line = strings.TrimSpace(line) // Remove leading/trailing spaces
		if logPattern.MatchString(line) {
			filtered = append(filtered, line)
		}
	}
	return filtered
}

func getLogFileSuffix(fileName string) int {
	if !strings.Contains(fileName, ".") {
		return 0
	}

	parts := strings.Split(fileName, ".")
	if len(parts) > 1 {
		if num, err := strconv.Atoi(parts[len(parts)-1]); err == nil {
			return num
		}
	}

	return 0
}

func GetLatestLogFile(dirListing string) (string, error) {
	files := filterLogFiles(dirListing)
	if len(files) == 0 {
		return "", fmt.Errorf("no mysql_error_log files found")
	}

	sort.Slice(files, func(i, j int) bool {
		iSuffix := getLogFileSuffix(files[i])
		jSuffix := getLogFileSuffix(files[j])
		return iSuffix > jSuffix
	})

	return files[0], nil
}
