package ini

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	ini := NewINI()
	err := ini.Parse("example.ini")
	if err != nil {
		fmt.Println("Error parsing INI file:", err)
		return
	}

	// Example usage
	value, exists := ini.Get("SectionName", "KeyName")
	if exists {
		fmt.Println("Value:", value)
	} else {
		fmt.Println("Key not found")
	}
}
