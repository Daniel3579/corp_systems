package main

import (
	"corp_systems/Practice_1/internal/FileData"
	"corp_systems/Practice_1/internal/WordData"
	"fmt"
)

func main() {
	var filePath string
	var wordToFind string

	fmt.Print("Введите путь к текстовому файлу: ")
	fmt.Scanln(&filePath)

	fmt.Print("Введите слово для поиска: ")
	fmt.Scanln(&wordToFind)

	var err error

	var file *FileData.Struct = &FileData.Struct{}
	file, err = file.Init(filePath)
	if err != nil {
		fmt.Println("Ошибка инициализации файла: ", err)
		return
	}

	var word *WordData.Struct = &WordData.Struct{}
	word, err = word.Init(file, wordToFind)
	if err != nil {
		fmt.Println("Ошибка инициализации слова: ", err)
		return
	}

	file.Print()
	word.Print()
}
