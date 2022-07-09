package repository

import (
	"bufio"
	"bytes"
	"os"
)

type FileTranslate struct {
	Path string
}

func NewFileTranslate(file string) *FileTranslate {
	return &FileTranslate{
		Path: file,
	}
}

func (ft *FileTranslate) GetWords() (string, error) {
	file, err := os.Open(ft.Path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	wr := bytes.Buffer{}
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		wr.WriteString(sc.Text())
	}
	return wr.String(), nil
}
