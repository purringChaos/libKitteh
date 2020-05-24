package filesystem

import (
	"os"
	"strconv"
)

func WriteBytes(fn string, bytes []byte) error {
	f, err := os.OpenFile(fn, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(bytes)
	if err != nil {
		return err
	}
	return err
}

func WriteString(fn string, data string) error {
	return WriteBytes(fn, []byte(data))
}

func WriteInt(fn string, data int) error {
	return WriteBytes(fn, []byte(strconv.Itoa(data)))
}