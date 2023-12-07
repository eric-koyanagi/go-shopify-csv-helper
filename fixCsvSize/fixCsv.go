package fixCsvSize

import (
	"encoding/csv"
	"fmt"
	"os"
)

const MAX_SIZE = 15 * 1024 * 1024
const MAX_ROWS_PER_FILE = 10000

func Check(file string) (bool, error) {
	return isValidSize(file)
}

func Split(fileName string, header []string, rows [][]string) error {
	processedRows := 0
	fileNumber := 1

	for processedRows < len(rows) {
		fileName := fmt.Sprintf("%s_%d.csv", fileName, fileNumber)
		file, err := os.Create(fileName)
		if err != nil {
			return err
		}
		defer file.Close()

		// Create a CSV writer and write the header
		writer := csv.NewWriter(file)
		err = writer.Write(header)
		if err != nil {
			return err
		}

		// Write up to MAX_ROWS_PER_FILE into the CSV
		for i := processedRows; i < len(rows) && i < processedRows+MAX_ROWS_PER_FILE; i++ {
			err = writer.Write(rows[i])
			if err != nil {
				return err
			}
		}

		fmt.Printf("Wrote CSV file batch #%d", processedRows)
		processedRows += MAX_ROWS_PER_FILE
		fileNumber++

		writer.Flush()
	}

	return nil
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
