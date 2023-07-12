package sugar

import (
	"fmt"
	"testing"
)

func TestMapEach(t *testing.T) {
	given := []string{"a", "b", "c"}
	expected := []string{"'a'", "'b'", "'c'"}
	result := MapEach(given, func(s string) string {
		return fmt.Sprintf("'%s'", s)
	})
	for i, item := range result {
		if item != expected[i] {
			t.Errorf("Expected %s, got %s", expected[i], item)
		}
	}
}
