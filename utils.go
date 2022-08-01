package RestoreAssets

import (
	"embed"
	"os"
	"path"
	"strings"
)

// Get current directory
func getCurrentDir() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return dir, nil
}

// Write a file to disk
func writeFile(dir string, data []byte) error {
	return os.WriteFile(dir, data, 0644)
}

// Create a new directory
func createDir(dir string) error {
	return os.Mkdir(dir, 0755)
}

// Return all dirnames available in the given FS
func getAllDirnames(fs *embed.FS, dir string) (out []string, err error) {
	if len(dir) == 0 {
		dir = "."
	}

	entries, err := fs.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		fp := path.Join(dir, entry.Name())
		if entry.IsDir() {
			out = append(out, fp)
			res, err := getAllDirnames(fs, fp)
			if err != nil {
				return nil, err
			}

			out = append(out, res...)

			continue
		}
	}

	return
}

// Return all filenames available in the given FS
func getAllFilenames(fs *embed.FS, dir string) (out []string, err error) {
	if len(dir) == 0 {
		dir = "."
	}

	entries, err := fs.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		fp := path.Join(dir, entry.Name())
		if entry.IsDir() {
			res, err := getAllFilenames(fs, fp)
			if err != nil {
				return nil, err
			}

			out = append(out, res...)

			continue
		}

		out = append(out, fp)
	}

	return
}

// Change parent directory name to new name
func changeDirName(filePath string, newDir string) string {
	if len(filePath) <= 0 {
		return ""
	}

	if newDir == "" {
		return filePath
	}

	if filePath[0] == '/' {
		filePath = filePath[1:]
	}

	index := strings.Index(filePath, "/")
	if index == -1 {
		return newDir
	}
	println()
	return strings.Replace(filePath, filePath[:index], newDir, 1)
}
