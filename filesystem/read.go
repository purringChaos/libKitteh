package filesystem

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadString(fn string) (string, error) {
	fileContents, err := ioutil.ReadFile(fn)
	if err != nil {
		return "", err
	}
	fileContentsString := string(fileContents)
	fileContentsString = strings.TrimSpace(fileContentsString)
	return fileContentsString, nil
}

func ReadFloat(fn string) (float64, error) {
	fileContentsString, err := ReadString(fn)
	if err != nil {
		return 0, err
	}
	parsedFloat, err := strconv.ParseFloat(fileContentsString, 64)
	if err != nil {
		return 0, err
	}
	return parsedFloat, nil
}

func ReadInt(fn string) (int64, error) {
	fileContentsString, err := ReadString(fn)
	if err != nil {
		return 0, err
	}
	parsedInt, err := strconv.ParseInt(fileContentsString, 10, 64)
	if err != nil {
		return 0, err
	}
	return parsedInt, nil
}

