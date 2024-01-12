package tools

import (
	"fmt"
	"testing"
)

func TestHideMiddlePhone(t *testing.T) {
	phoneNumber := "12345678901"
	hidden := HideMiddlePhone(phoneNumber)
	fmt.Println(hidden)
}
