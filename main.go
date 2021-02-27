package main

import (
	"embed"
	_ "embed"
	"fmt"
)

//go:embed files
var f embed.FS

func main() {
	fmt.Println("Playing with embedded files")
	data, _ := f.ReadFile("files/a.txt")
	fmt.Println(string(data))
	data, _ = f.ReadFile("files/b.txt")
	fmt.Println(string(data))

	dirEntries, _ := f.ReadDir("files")
	for _, dirEntry := range dirEntries {
		fmt.Println(dirEntry.Name())
	}
}
