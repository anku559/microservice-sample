package helpers

import (
	"os"
)

func ReadTextFile(filename string) (string, error) {
	content, err := os.ReadFile(filename)

	if err != nil {
		return "", err
	}
	return string(content), nil

}