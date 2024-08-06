package routers

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

func ParseLocationCodes(filePath string) (map[string]bool, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	locationCodes := make(map[string]bool)
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`'([A-Z0-9]+)':\s*'`)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, ":") {
			matches := re.FindStringSubmatch(line)
			if len(matches) > 1 {
				locationCodes[matches[1]] = true
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return locationCodes, nil
}
