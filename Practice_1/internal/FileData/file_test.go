package FileData

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadContent(t *testing.T) {
	var cases = []struct {
		name string
		path string
		want string
		err  bool
	}{
		{"Wrong path", "test0.txt", "", true},
		{"No Line", "../../test_files/No Line.txt", "", false},
		{"One Line", "../../test_files/One Line.txt", "line 1\n", false},
		{"Two Lines", "../../test_files/Two Lines.txt", "line 1\nline 2\n", false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := readContent(c.path)

			if c.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, c.want, got)
		})
	}
}

func TestParsingName(t *testing.T) {
	var errNotFound error = errors.New("Имя файла не найдено")

	var cases = []struct {
		path string
		want string
		err  error
	}{
		{"txt", "txt", nil},
		{"txt.", "txt", nil},
		{".txt", "", errNotFound},
		{"file.txt", "file", nil},
		{"file.ext.txt", "file.ext", nil},
		{"/folder/folder/txt", "txt", nil},
		{"/folder/folder/txt.", "txt", nil},
		{"/folder/folder/.txt", "", errNotFound},
		{"/folder/folder/file.txt", "file", nil},
		{"/folder/folder/file.ext.txt", "file.ext", nil},
	}

	for _, c := range cases {
		t.Run(c.path, func(t *testing.T) {
			got, err := parsingName(c.path)

			if c.err != nil {
				assert.EqualError(t, err, errNotFound.Error())
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, c.want, got)
		})
	}
}

func TestParsingExtension(t *testing.T) {
	var errNotFound = errors.New("Расширение файла не найдено")

	var cases = []struct {
		path string
		want string
		err  error
	}{
		{"txt", "", errNotFound},
		{"txt.", ".", nil},
		{".txt", ".txt", nil},
		{"file.txt", ".txt", nil},
		{"file.ext.txt", ".txt", nil},
		{"/folder/folder/txt", "", errNotFound},
		{"/folder/folder/txt.", ".", nil},
		{"/folder/folder/.txt", ".txt", nil},
		{"/folder/folder/file.txt", ".txt", nil},
		{"/folder/folder/file.ext.txt", ".txt", nil},
	}

	for _, c := range cases {
		t.Run(c.path, func(t *testing.T) {
			got, err := parsingExtension(c.path)

			if err != nil {
				assert.EqualError(t, err, errNotFound.Error())
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, c.want, got)
		})
	}
}

func TestCountWords(t *testing.T) {
	var cases = []struct {
		content string
		want    int
	}{
		{"", 0},
		{"           ", 0},
		{"Word", 1},
		{"Word\n", 1},
		{"Word ", 1},
		{"Word \n", 1},
		{" Word", 1},
		{" Word ", 1},
		{"    Word    ", 1},
		{"One two", 2},
		{"o n e t w o", 6},
	}

	for _, c := range cases {
		t.Run(c.content, func(t *testing.T) {
			var got int = countWords(c.content)
			assert.Equal(t, c.want, got)
		})
	}
}
