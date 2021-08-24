package main

import (
	"github.com/stretchr/testify/require"
	"testing"
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
}
