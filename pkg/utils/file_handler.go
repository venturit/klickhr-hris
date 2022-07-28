package utils

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
)

var correctHeaders = []string{"First Name",
	"Last Name",
	"Employee ID",
	"Emplpyee Status",
	"Email Address",
	"Hire Date",
	"Start Date",
	"Status Change Date",
	"Anniversary M/D",
	"Supervisor 1",
	"Supervisor 2",
	"Supervisor 3 ",
	"Position Type",
	"Level 1 Code",
	"Level 2 Code",
	"Level 3 Code",
	"Level 4 Code",
	"Street Address",
	"City",
	"State",
	"Zip Code",
	"FLSA Status",
	"Union"}

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

func ValidateFile(path string) error {
	//extension validation
	fileExtension := filepath.Ext(path)
	if fileExtension != ".csv" {
		return errors.New("File extension ins't equal to .csv")
	}
	//header validation
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	r := csv.NewReader(f)
	headers, err := r.Read()
	if err != nil {
		return err
	}
	var stringHeaders string
	for _, header := range headers {
		stringHeaders += header + ", "
	}
	var stringCorrectHeaders string
	for _, correctheader := range correctHeaders {
		stringCorrectHeaders += correctheader + ", "
	}
	fmt.Println(stringHeaders)
	if !(reflect.DeepEqual(headers, correctHeaders)) {
		return errors.New("the headers delivered are not the corresponding ones:\n" + stringCorrectHeaders + "\nwas obtained:\n" + stringHeaders)
	}
	return nil
}
