package WordData

import (
	"corp_systems/Practice_1/internal/FileData"
	"errors"
	"fmt"
	"strings"
)

type Struct struct {
	file  *FileData.Struct
	word  string
	count int
}

func (wordData *Struct) Init(file *FileData.Struct, word string) (*Struct, error) {
	if file == nil {
		return nil, errors.New("Нет ссылки на файл")
	}

	wordData.file = file
	wordData.word = word
	wordData.count = wordData.countWord()

	return wordData, nil
}

// Функция для поиска заданного слова и подсчета его повторений
func (wordData *Struct) countWord() int {
	words := strings.Fields(wordData.file.GetContent())
	count := 0

	for _, w := range words {
		if strings.EqualFold(w, wordData.word) {
			count++
		}
	}

	return count
}

func (wordData *Struct) Print() {
	fmt.Printf("\n"+
		"File: %p\n"+
		"Word: %s\n"+
		"Count: %d\n"+
		"\n",
		wordData.file,
		wordData.word,
		wordData.count,
	)
}
