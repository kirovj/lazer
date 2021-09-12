package lazer

import "os"

type FileWriter struct {
	File     *os.File
	Filename string
}

func NewFileWriter(filename string) *FileWriter {
	return &FileWriter{
		Filename: filename,
	}
}

func (f *FileWriter) Ready() bool {
	file, err := os.OpenFile(f.Filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return false
	}
	f.File = file
	return true
}

func (f *FileWriter) Write(content []byte) error {
	if _, err := f.File.Write(content); err != nil {
		return err
	}
	return nil
}

func (f *FileWriter) Close() error {
	if err := f.File.Close(); err != nil {
		return err
	}
	return nil
}
