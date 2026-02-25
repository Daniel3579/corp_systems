package FileData

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Struct struct {
	content   string
	name      string
	extension string
	wordCount int
}

func (fileData *Struct) Init(path string) (*Struct, error) {
	content, err := readContent(path)
	if err != nil {
		return nil, err
	}

	name, err := parsingName(path)
	if err != nil {
		fileData = nil
		return nil, err
	}

	extension, err := parsingExtension(path)
	if err != nil {
		fileData = nil
		return nil, err
	}

	var wordCount int = countWords(content)

	fileData.content = content
	fileData.name = name
	fileData.extension = extension
	fileData.wordCount = wordCount

	return fileData, nil
}

func readContent(path string) (string, error) {
	// Попытка открытия файла
	file, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("Ошибка при открытии файла \"%s\": %w", path, err)
	}
	defer file.Close()

	// Чтение файла
	var content strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content.WriteString(scanner.Text() + "\n")
	}

	return content.String(), nil
}

func parsingName(path string) (string, error) {
	var name string = strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
	if name == "" {
		return "", errors.New("Имя файла не найдено")
	}
	return name, nil
}

func parsingExtension(path string) (string, error) {
	var extension string = filepath.Ext(path)
	if extension == "" {
		return "", errors.New("Расширение файла не найдено")
	}
	return extension, nil
}

func countWords(content string) int {
	var words []string = strings.Fields(content)
	return len(words)
}

func (fileData *Struct) GetContent() string {
	return fileData.content
}

func (fileData *Struct) GetName() string {
	return fileData.name
}

func (fileData *Struct) GetExtension() string {
	return fileData.extension
}

func (fileData *Struct) GetWordCount() int {
	return fileData.wordCount
}

func (fileData *Struct) Print() {
	fmt.Printf("\n"+
		"Content: %s\n"+
		"Name: %s\n"+
		"Extension: %s\n"+
		"WordCount: %d\n"+
		"\n",
		fileData.content,
		fileData.name,
		fileData.extension,
		fileData.wordCount,
	)
}
