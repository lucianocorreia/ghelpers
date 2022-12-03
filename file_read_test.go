package ghhelpers

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadlinesFromReader_ShouldReadTheBuffer(t *testing.T) {
	var buffer bytes.Buffer
	buffer.WriteString("line1\nline2\nline3\n")

	lines, err := readlinesFromReader(&buffer)
	require.NoError(t, err)

	require.Equal(t, 3, len(lines))
	require.Equal(t, lines[0], "line1")
	require.Equal(t, lines[1], "line2")
	require.Equal(t, lines[2], "line3")
}

func TestFileReadLines_ShouldReadTheFile(t *testing.T) {
	lines, err := FileReadLines("test.txt")
	require.NoError(t, err)
	require.NotEmpty(t, lines)
	require.Equal(t, "line1", lines[0])
}

func TestFileREadline_ShouldReturnErroIfFileNotFound(t *testing.T) {
	lines, err := FileReadLines("file.notexists")
	require.Error(t, err)
	require.Empty(t, lines)
}
