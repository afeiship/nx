package nx

import (
	"io/ioutil"
	"os"
)

func WriteFile(fileName string, data []byte) error {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func ReadFile(fileName string) ([]byte, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func Exists(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil
}

func Remove(fileName string) error {
	return os.Remove(fileName)
}

func Rename(oldName, newName string) error {
	return os.Rename(oldName, newName)
}

func Mkdir(dirName string) error {
	return os.Mkdir(dirName, 0755)
}

func ReadDir(dirName string) ([]os.FileInfo, error) {
	return ioutil.ReadDir(dirName)
}

func IsFile(fileName string) bool {
	fi, err := os.Stat(fileName)
	if err != nil {
		return false
	}
	return !fi.IsDir()
}

func IsDir(fileName string) bool {
	fi, err := os.Stat(fileName)
	if err != nil {
		return false
	}
	return fi.IsDir()
}
