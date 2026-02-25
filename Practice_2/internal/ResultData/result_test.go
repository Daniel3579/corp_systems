package ResultData

import (
	"sync"
	"testing"
)

func TestConcurrentAddWords(t *testing.T) {
	data := &Struct{}
	var wg sync.WaitGroup
	numGoroutines := 100

	for i := 1; i <= numGoroutines; i++ {
		wg.Add(1)

		go func(count int) {
			defer wg.Done()
			data.AddWords(count)
		}(i)
	}

	wg.Wait()

	expectedWordCount := (numGoroutines * (numGoroutines + 1)) / 2

	if data.GetWordCount() != expectedWordCount {
		t.Errorf("Expected word count %d, got %d", expectedWordCount, data.GetWordCount())
	}
}

func TestConcurrentAddSymbols(t *testing.T) {
	data := &Struct{}
	var wg sync.WaitGroup
	numGoroutines := 100

	for i := 1; i <= numGoroutines; i++ {
		wg.Add(1)

		go func(count int) {
			defer wg.Done()
			data.AddSymbols(count * 2)
		}(i)
	}

	wg.Wait()

	expectedSymbolCount := (numGoroutines * (numGoroutines + 1))

	if data.GetSymbolCount() != expectedSymbolCount {
		t.Errorf("Expected symbol count %d, got %d", expectedSymbolCount, data.GetSymbolCount())
	}
}
