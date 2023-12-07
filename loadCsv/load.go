package loadCsv

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"

	"csvHelper/validateHeader"
)

var ErrLoadCSV = errors.New("failed to load CSV")

func GetCSV(path string) ([][]string, error) {
	// Open the CSV file
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("opening file %w", err)
	}
	defer file.Close()

	// Load the CSV data
	data, err := loadCSV(file)
	if err != nil {
		return nil, err
	}

	// Check that the header is valid
	isValid, err := validateHeader.Validate(data[0])
	if !isValid {
		return nil, fmt.Errorf("column header is not valid: %w", err)
	}

	return data, err
}

func loadCSV(r io.Reader) ([][]string, error) {
	reader := csv.NewReader(r)

	// Read all records
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("reading CSV: %w", err)
	}

	return records, nil
}
