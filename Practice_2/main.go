package main

import (
	"corp_systems/Practice_2/internal/FileAnalysis"
	"corp_systems/Practice_2/internal/ResultData"
	"fmt"
	"sync"
)

var res ResultData.Struct = ResultData.Struct{}
var wg sync.WaitGroup

func main() {
	var paths [8]string = [8]string{
		"./test_files/test1.txt",
		"./test_files/test2.txt",
		"./test_files/test3.txt",
		"./test_files/test4.txt",
		"./test_files/test5.txt",
		"./test_files/test6.txt",
		"./test_files/test7.txt",
		"./test_files/test8.txt",
	}

	wg.Add(len(paths))
	for _, path := range paths {
		go do(&wg, path)
	}
	wg.Wait()

	res.Print()
}

func do(wg *sync.WaitGroup, path string) {
	var file *FileAnalysis.Struct = &FileAnalysis.Struct{}
	file, err := file.Init(path)
	if err != nil {
		fmt.Println("Ошибка инициализации файла: ", err)
		return
	}

	res.Add(*file)
	wg.Done()
}
