package main

import (
	"embed"
	_ "embed"
	"fmt"
	"github.com/pkg/errors"
	"io/fs"
	"log"
	"os"
)

//go:embed files
var f embed.FS

type readable interface {
	fs.ReadDirFS
	fs.ReadFileFS
}

type localFS string

func (l localFS) Open(name string) (fs.File, error)          { return os.Open(name) }
func (l localFS) ReadDir(name string) ([]fs.DirEntry, error) { return os.ReadDir(name) }
func (l localFS) ReadFile(name string) ([]byte, error)       { return os.ReadFile(name) }

func main() {
	fmt.Println("Playing with embedded files")

	fmt.Println("From local fs")
	dumpFiles(localFS("."), "files")

	fmt.Println("***************************")

	fmt.Println("From embedded fs")
	dumpFiles(f, "files")

}

func dumpFiles(fs readable, dirname string) {
	dirEntries, err := fs.ReadDir(dirname)
	if err != nil {
		log.Fatal(errors.Wrap(err, "had issues listing the files"))
	}
	for _, dirEntry := range dirEntries {
		if dirEntry.IsDir() {
			fmt.Printf(">>> %s\n", dirEntry.Name())
			dumpFiles(fs, dirname+"/"+dirEntry.Name())
			fmt.Println("<<<")
		} else {
			fmt.Printf("==== %s ====\n", dirEntry.Name())
			data, _ := fs.ReadFile(dirname + "/" + dirEntry.Name())
			fmt.Println(string(data))
			fmt.Println("==== END ====")
		}
	}
}
