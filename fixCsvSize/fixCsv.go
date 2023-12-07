package fixCsvSize

import (
	"os"
)

const MAX_SIZE = 15 * 1024 * 1024

func Check(file string) (bool, error) {
	return isValidSize(file)
}

func Split(header []string, rows [][]string) {

}

func isValidSize(path string) (bool, error) {
	// Open the file
	file, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer file.Close()

	// Get file size
	fileInfo, err := file.Stat()
	if err != nil {
		return false, err
	}

	// Check if size is less than 15 MB (15 * 1024 * 1024)
	return fileInfo.Size() < MAX_SIZE, nil
}
