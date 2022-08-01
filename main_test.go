package RestoreAssets

import (
	"embed"
	"os"
	"testing"
)

//go:embed test/*
var assets embed.FS

func TestMain(t *testing.T) {
	// RestoreAssets.From(&fs, "dirToWrite")
	err := From(&assets, "test_out")
	if err != nil {
		t.Error(err)
	}

	res, err := os.ReadFile("test_out/example.txt")
	if err != nil {
		t.Error(err)
	}

	if string(res) != "Example file used for testing purposes only." {
		t.Error("File content does not match")
	}
	os.RemoveAll("test_out")
}
