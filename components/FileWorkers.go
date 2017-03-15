package components

import (
	"bufio"
	"errors"
	"io"
	"os"
	"path/filepath"
)

func checkFileExist(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func createDir(dirPath string, fileMode os.FileMode) bool {
	err := os.MkdirAll(dirPath, fileMode)
	if err != nil {
		return false
	}
	return true
}

func renameFile(pathFile string, name string) error {
	err := os.Rename(pathFile, name)
	if err != nil {
		return err
	}
	return nil
}

func createFile(pathFile string) error {
	file, err := os.Create(pathFile)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

func createFileAndWriteData(fileName string, writeData []byte) error {
	fileHandle, err := os.Create(fileName)

	if err != nil {
		return err
	}
	writer := bufio.NewWriter(fileHandle)
	defer fileHandle.Close()
	writer.Write(writeData)
	writer.Flush()
	return nil
}

func copyFileToDirectory(pathSourceFile string, pathDestFile string) error {
	sourceFile, err := os.Open(pathSourceFile)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(pathDestFile)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	err = destFile.Sync()
	if err != nil {
		return err
	}

	sourceFileInfo, err := sourceFile.Stat()
	if err != nil {
		return err
	}

	destFileInfo, err := destFile.Stat()
	if err != nil {
		return err
	}

	if sourceFileInfo.Size() == destFileInfo.Size() {
	} else {
		err = errors.New("Bad copy file")
		return err
	}
	return nil
}

func deleteFile(nameFile string) error {
	err := os.Remove(nameFile)
	return err
}

func removeDirWithContent(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	err = os.RemoveAll(dir)
	if err != nil {
		return err
	}
	return nil
}
