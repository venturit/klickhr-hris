package utils

import (
	"encoding/csv"
	"fmt"
	"os"
)

func SaveFile(fileName string, fileBytes []byte) error {
	err := os.WriteFile(fmt.Sprintf("./%s", fileName), fileBytes, 0644)
	if err != nil {
		return err
	}
	return nil
}

func ReadCSVData(fileName string) ([][]string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()
	r := csv.NewReader(f)
	// skip first line
	if _, err := r.Read(); err != nil {
		return [][]string{}, err
	}
	records, err := r.ReadAll()
	if err != nil {
		return [][]string{}, err
	}
	return records, nil
}
