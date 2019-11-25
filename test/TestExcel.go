package main

import (
	"encoding/csv"
	"os"

	"github.com/cuijxin/go-gin-example/pkg/export"
)

func main() {
	f, err := os.Create(export.GetExcelFullPath() + "test.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("\xFF\xBB\xBF")

	w := csv.NewWriter(f)
	data := [][]string{
		{"1", "test1", "test1-1"},
		{"2", "test2", "test2-1"},
		{"3", "test3", "test3-1"},
	}

	w.WriteAll(data)
}
