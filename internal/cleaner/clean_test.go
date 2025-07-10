package cleaner

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestCleanUpOldFilesDryRun(t *testing.T) {
	tempDir := t.TempDir()

	oldFile := filepath.Join(tempDir, "old.txt")
	newFile := filepath.Join(tempDir, "new.txt")

	os.WriteFile(oldFile, []byte("unalive me"), 0644)
	os.WriteFile(newFile, []byte("I'm staying where I am"), 0644)

	oldTime := time.Now().AddDate(0, 0, -40)
	err := os.Chtimes(oldFile, oldTime, oldTime)
	if err != nil {
		t.Fatalf("Failed to set old mod time: %v", err)
	}

	err = CleanUp(tempDir, 30, true)
	if err != nil {
		t.Fatalf("Dry run failed: %v", err)
	}

	if _, err := os.Stat(oldFile); os.IsNotExist(err) {
		t.Fatalf("Old file was deleted during dry run: %v", err)
	}

	if _, err := os.Stat(newFile); os.IsNotExist(err) {
		t.Fatalf("New file was deleted during dry run: %v", err)
	}
}

func TestCleanUpOldFiles(t *testing.T) {
	tempDir := t.TempDir()

	oldFile := filepath.Join(tempDir, "old.txt")
	newFile := filepath.Join(tempDir, "new.txt")

	os.WriteFile(oldFile, []byte("unalive me"), 0644)
	os.WriteFile(newFile, []byte("I'm staying where I am"), 0644)

	oldTime := time.Now().AddDate(0, 0, -40)
	err := os.Chtimes(oldFile, oldTime, oldTime)
	if err != nil {
		t.Fatalf("Failed to set old mod time: %v", err)
	}

	err = CleanUp(tempDir, 30, false)
	if err != nil {
		t.Fatalf("Clean up failed: %v", err)
	}

	if _, err := os.Stat(oldFile); !os.IsNotExist(err) {
		t.Fatalf("Old file was not deleted: %v", err)
	}

	if _, err := os.Stat(newFile); os.IsNotExist(err) {
		t.Fatalf("New file was deleted: %v", err)
	}
}

func TestEmptyDirCleanUp(t *testing.T) {
	tempDir := t.TempDir()

	err := CleanUp(tempDir, 30, true)
	if err != nil {
		t.Fatalf("Dry run failed: %v", err)
	}

	err = CleanUp(tempDir, 30, false)
	if err != nil {
		t.Fatalf("Clean up failed: %v", err)
	}
}

func TestCleanUpOldFilesNestedDirectories(t *testing.T) {
	tempDir := t.TempDir()

	files := []*struct {
		filename string
		path     string
		old      bool
		absPath  string
	}{
		{
			filename: "test1.json",
			path:     "a/b/c",
			old:      true,
		},
		{
			filename: "test2.json",
			old:      false,
		},
		{
			filename: "test3.md",
			path:     "d/e/c",
			old:      true,
		},
		{
			filename: "test4.json",
			path:     "b/c/a/b/c/a/b/c",
			old:      false,
		},
		{
			filename: "test5.go",
			path:     "m/n",
			old:      true,
		},
		{
			filename: "test6.txt",
			path:     "a/b/c",
			old:      false,
		},
		{
			filename: "test7.go",
			old:      true,
		},
		{
			filename: "test8.txt",
			old:      false,
		},
	}

	for _, file := range files {
		if file.path != "" {
			file.absPath = filepath.Join(tempDir, file.path, file.filename)
			err := os.MkdirAll(filepath.Dir(file.absPath), 0755)
			if err != nil {
				t.Fatalf("Failed to create parent directories %s: %v", file.path, err)
			}
		} else {
			file.absPath = filepath.Join(tempDir, file.filename)
		}

		err := os.WriteFile(file.absPath, []byte("testing "+file.filename), 0644)
		if err != nil {
			t.Fatalf("Failed to write to file %s: %v", file.filename, err)
		}

		if file.old {
			oldTime := time.Now().AddDate(0, 0, -40)
			if err := os.Chtimes(file.absPath, oldTime, oldTime); err != nil {
				t.Fatalf("Failed to set old mod time %s: %v", file.filename, err)
			}
		}

	}

	err := CleanUp(tempDir, 30, false)
	if err != nil {
		t.Fatalf("Clean up failed: %v", err)
	}

	for _, file := range files {
		_, err := os.Stat(file.absPath)

		if file.old && !os.IsNotExist(err) {
			t.Fatalf("Old file was not deleted %s: %v", file.filename, err)
		}

		if !file.old && os.IsNotExist(err) {
			t.Fatalf("New file was deleted %s: %v", file.filename, err)
		}
	}
}
