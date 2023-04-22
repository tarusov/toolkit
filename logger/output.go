package logger

import "os"

// OpenFile - open file with append parameters.
func OpenFile(filename string) (file *os.File, err error) {
	file, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	return file, err
}
