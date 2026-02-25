package FileAnalysis

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Struct struct {
	name        string
	wordCount   int
	symbolCount int
}

func (fileAnalysis *Struct) Init(path string) (*Struct, error) {
	content, err := readContent(path)
	if err != nil {
		return nil, err
	}

	name, err := parsingName(path)
	if err != nil {
		fileAnalysis = nil
		return nil, err
	}

	var wordCount int = countWords(content)
	var symbolCount int = countSymbols(content)

	fileAnalysis.name = name
	fileAnalysis.wordCount = wordCount
	fileAnalysis.symbolCount = symbolCount

	return fileAnalysis, nil
}

func readContent(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("Ошибка при открытии файла \"%s\": %w", path, err)
	}
	defer file.Close()

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
		return "", fmt.Errorf("Имя файла не найдено")
	}
	return name, nil
}

func countWords(content string) int {
	var words []string = strings.Fields(content)
	return len(words)
}

func countSymbols(content string) int {
	return len(content)
}

func (fileAnalysis *Struct) GetName() string {
	return fileAnalysis.name
}

func (fileAnalysis *Struct) GetWordCount() int {
	return fileAnalysis.wordCount
}

func (fileAnalysis *Struct) GetSymbolCount() int {
	return fileAnalysis.symbolCount
}

func (fileAnalysis *Struct) Print() {
	fmt.Printf("\n"+
		"Name: %s\n"+
		"WordCount: %d\n"+
		"SymbolCount: %d\n"+
		"\n",
		fileAnalysis.name,
		fileAnalysis.wordCount,
		fileAnalysis.symbolCount,
	)
}
