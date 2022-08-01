package RestoreAssets

import (
	"embed"
	"os"
)

func From(fs *embed.FS, outDir string) error {
	dirs, err := getAllDirnames(fs, "")
	if err != nil {
		return err
	}

	currentDir, err := getCurrentDir()
	if err != nil {
		return err
	}

	// Create all directories (to prevent file creation errors)
	for _, dir := range dirs {
		err := createDir(currentDir + "\\" + changeDirName(dir, outDir))
		if err != nil {
			if os.IsExist(err) {
				continue
			}
			return err
		}
	}

	files, err := getAllFilenames(fs, "")
	if err != nil {
		return err
	}

	for _, file := range files {
		data, _ := fs.ReadFile(file)
		err := writeFile(currentDir+"/"+changeDirName(file, outDir), data)
		if err != nil {
			if os.IsExist(err) {
				continue
			}
			return err
		}
	}

	// no error
	return nil
}
