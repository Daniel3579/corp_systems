package ResultData

import (
	"corp_systems/Practice_2/internal/FileAnalysis"
	"fmt"
	"sync"
)

type Struct struct {
	wordCount   int
	symbolCount int
}

var mu sync.Mutex

func (data *Struct) Add(file FileAnalysis.Struct) {
	mu.Lock()
	defer mu.Unlock()

	data.wordCount += file.GetWordCount()
	data.symbolCount += file.GetSymbolCount()
}

func (data *Struct) AddWords(count int) {
	mu.Lock()
	defer mu.Unlock()

	data.wordCount += count
}

func (data *Struct) AddSymbols(count int) {
	mu.Lock()
	defer mu.Unlock()

	data.symbolCount += count
}

func (data *Struct) GetWordCount() int {
	return data.wordCount
}

func (data *Struct) GetSymbolCount() int {
	return data.symbolCount
}

func (data *Struct) Print() {
	fmt.Printf("\n"+
		"WordCount: %d\n"+
		"SymbolCount: %d\n"+
		"\n",
		data.wordCount,
		data.symbolCount,
	)
}
