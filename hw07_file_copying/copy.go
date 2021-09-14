package main

import (
	"errors"
	"fmt"
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
		return fmt.Errorf("error from Open: %w", err)
	}
	defer closeFile(file)

	fileStat, err := file.Stat()
	if err != nil {
		return fmt.Errorf("error from Stat: %w", err)
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
	if limit == 0 || limit > fileStat.Size()-offset {
		limit = fileStat.Size() - offset
	}

	var r io.Reader = file
	var buf []byte
	var bar *pb.ProgressBar

	if offset != 0 || limit != 0 {
		r = io.NewSectionReader(file, offset, limit)
		buf = make([]byte, io.NewSectionReader(file, offset, limit).Size())
		bar = pb.ProgressBarTemplate(tmpl).Start64(io.NewSectionReader(file, offset, limit).Size())
	} else {
		buf = make([]byte, fileStat.Size())
		bar = pb.ProgressBarTemplate(tmpl).Start64(fileStat.Size())
	}

	defer bar.Finish()

	if _, err = bar.NewProxyReader(r).Read(buf); err != nil && !errors.Is(err, io.EOF) {
		return fmt.Errorf("error from Read: %w", err)
	}
	if err = ioutil.WriteFile(toPath, buf, 0600); err != nil {
		return fmt.Errorf("error from WriteFile method: %w", err)
	}

	return nil
}
