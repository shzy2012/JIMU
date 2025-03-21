package ini

import (
	"bufio"
	"os"
	"strings"
)

// INI represents the parsed INI file structure
type INI struct {
	Sections map[string]map[string]string
}

// NewINI creates a new INI structure
func NewINI() *INI {
	return &INI{Sections: make(map[string]map[string]string)}
}

// Parse parses the INI file from the given filename
func (ini *INI) Parse(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var currentSection string

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}

		// Handle section headers
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			currentSection = strings.TrimSpace(line[1 : len(line)-1])
			if _, exists := ini.Sections[currentSection]; !exists {
				ini.Sections[currentSection] = make(map[string]string)
			}
			continue
		}

		// Handle key-value pairs
		if currentSection != "" {
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				key := strings.TrimSpace(parts[0])
				value := strings.TrimSpace(parts[1])
				ini.Sections[currentSection][key] = value
			}
		}
	}

	return scanner.Err()
}

// Get retrieves a value from the INI file given a section and key
func (ini *INI) Get(section, key string) (string, bool) {
	if sec, exists := ini.Sections[section]; exists {
		value, exists := sec[key]
		return value, exists
	}
	return "", false
}
