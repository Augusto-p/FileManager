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
	
		// Lee el contenido de la carpeta de origen
		entries, err := os.ReadDir(sourcePath)
		if err != nil {
			return err
		}
	
		// Recorre los archivos y subdirectorios en la carpeta de origen
		for _, entry := range entries {
			sourceFile := filepath.Join(sourcePath, entry.Name())
			destinationFile := filepath.Join(destinationPath, entry.Name())
	
			err = Copy(sourceFile, destinationFile)
			if err != nil {
        
				return err
			}
		}
	}else{
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
	return nil

}		

