package FileManager

import (
	"io"
	"os"
	"path/filepath"
)
func Copy(sourcePath, destinationPath string) error {
	fileInfo, err := os.Stat(sourcePath)

    if err != nil {
        
        return err
    }

    if fileInfo.IsDir() {
		if err := os.MkdirAll(destinationPath, os.ModePerm); err != nil {
			return err
		}
		entries, err := os.ReadDir(sourcePath)
		if err != nil {
			return err
		}
		for _, entry := range entries {
			sourceFile := filepath.Join(sourcePath, entry.Name())
			destinationFile := filepath.Join(destinationPath, entry.Name())
			Copy(sourceFile, destinationFile)
		}
	
    } else {
        sourceFile, err := os.Open(sourcePath)
		if err != nil {
		return err
		}
		defer sourceFile.Close()

		destinationFile, err := os.Create(destinationPath)
		if err != nil {
			return err
		}
		defer destinationFile.Close()

		_, err = io.Copy(destinationFile, sourceFile)
		if err != nil {
			return err
		}

		return nil
	}
}
