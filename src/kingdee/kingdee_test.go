package kingdee

import (
	"testing"
)

func TestKingdee(t *testing.T) {
	err := KingdeeGetInStock()
	if err != nil {
		t.Fatal(err)
	}
}
