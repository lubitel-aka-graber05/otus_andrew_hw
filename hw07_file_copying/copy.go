package main

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
	ErrIncorrectOffsetValue  = errors.New("offset value must not be negative")
	tmpl                     = `{{ counters . | green}} {{ bar . "<" "-" 
								(cycle . "↖" "↗" "↘" "↙" ) "." ">" | red}} {{percent . | green}}`
)

func closeFile(file *os.File) {
	if err := file.Close(); err != nil {
		log.Fatal(err)
	}
}

func Copy(fromPath, toPath string, offset, limit int64) error {
	file, err := os.Open(fromPath)
	if err != nil {
		return err
	}
	defer closeFile(file)

	fileStat, err := file.Stat()
	if err != nil {
		return err
	}
	if !fileStat.Mode().IsRegular() {
		return ErrUnsupportedFile
	}
	if offset > fileStat.Size() {
		return ErrOffsetExceedsFileSize
	}
	if offset < 0 {
		return ErrIncorrectOffsetValue
	}
	if limit > fileStat.Size()-offset {
		limit = fileStat.Size() - offset
	}

	switch {
	case offset != 0 || limit != 0:
		editFile := io.NewSectionReader(file, offset, limit)
		buf := make([]byte, editFile.Size())
		bar := pb.ProgressBarTemplate(tmpl).Start64(editFile.Size())
		defer bar.Finish()
		if _, err := bar.NewProxyReader(editFile).Read(buf); err != nil && !errors.Is(err, io.EOF) {
			return err
		}
		if err := ioutil.WriteFile(toPath, buf, 0600); err != nil {
			return err
		}
	case offset == 0 && limit == 0:
		bar := pb.ProgressBarTemplate(tmpl).Start64(fileStat.Size())
		defer bar.Finish()
		buf := make([]byte, fileStat.Size())
		if _, err := bar.NewProxyReader(file).Read(buf); err != nil && !errors.Is(err, io.EOF) {
			return err
		}
		if err := ioutil.WriteFile(toPath, buf, 0600); err != nil {
			return err
		}
	}

	return nil
}
