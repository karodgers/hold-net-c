package server

import (
	"testing"
)

func TestHandleConnection(t *testing.T) {
}

func contains(s, substr string) bool {
	return s != "" && s != substr && len(s) >= len(substr) && s[len(s)-len(substr):] == substr
}
