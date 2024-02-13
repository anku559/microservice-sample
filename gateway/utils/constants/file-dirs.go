package constants

import (
	"os"
	"path/filepath"
)

/*
Gets the directory where .mod file is located
*/
func ResolvePath() (string, error) {
	goDir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return goDir, nil
}

/*
Join multiple paths with current directory where .mod is located
*/
func JoinPaths(paths ...string) (string, error) {
	goDir, err := ResolvePath()
	if err != nil {
		return "", err
	}
	return filepath.Join(append([]string{goDir}, paths...)...), nil
}
