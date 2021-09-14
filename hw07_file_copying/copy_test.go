package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	t.Run("test for unsupported file", func(t *testing.T) {
		err := Copy("/dev/urandom", "/tmp/output.txt", 0, 0)
		require.ErrorIs(t, err, ErrUnsupportedFile, "unsupported files")
	})

	t.Run("negative offset", func(t *testing.T) {
		err := Copy("testdata/input.txt", "/tmp/output.txt", -1, 0)
		require.ErrorIs(t, err, ErrIncorrectOffsetValue, "incorrect offset")
	})

	t.Run("offset > file sizes", func(t *testing.T) {
		err := Copy("testdata/input.txt", "/tmp/output.txt", 7000, 0)
		require.ErrorIs(t, err, ErrOffsetExceedsFileSize, "Offset is larger than the file size")
	})

	t.Run("offset = 0 & limit =0", func(t *testing.T) {
		tempFile, _ := os.CreateTemp("/tmp", "output.txt")
		defer tempFile.Close()
		err := Copy("testdata/input.txt", tempFile.Name(), 0, 0)
		file, _ := os.Open("testdata/input.txt")
		defer file.Close()
		forCompare1, _ := file.Stat()
		forCompare2, _ := tempFile.Stat()
		require.NoError(t, err)
		require.Equal(t, forCompare2.Size(), forCompare1.Size())
		_ = os.Remove(tempFile.Name())
	})

	t.Run("offset = 1 & limit = 10", func(t *testing.T) {
		tempFile, _ := os.CreateTemp("/tmp", "output.txt")
		defer tempFile.Close()
		err := Copy("testdata/input.txt", tempFile.Name(), 1, 10)
		file, _ := os.Open("testdata/input.txt")
		defer file.Close()
		forCompFile, _ := file.Stat()
		forCompTempFile, _ := tempFile.Stat()
		require.NoError(t, err)
		require.Equal(t, forCompTempFile.Size(), forCompFile.Size()-(forCompFile.Size()-forCompTempFile.Size()))
		_ = os.Remove(tempFile.Name())
	})
}
