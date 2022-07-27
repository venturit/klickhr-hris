package utils

import (
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
