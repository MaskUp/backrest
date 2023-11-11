package test

import (
	"fmt"
	"os"
	"path"
	"testing"
)

func CreateTestData(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	
	for i := 0; i < 100; i++ {
		err := os.WriteFile(path.Join(dir, fmt.Sprintf("file%2d", i)), []byte(fmt.Sprintf("test data %d", i)), 0644)
		if err != nil {
			t.Fatalf("failed to create test data: %v", err)
		}
	}
	return dir
}