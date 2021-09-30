package learn_go_embed

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

// https://pkg.go.dev/embed
// "go:embed [name of the file]" to load/embed file in golang
// load it into a variable outside function
//go:embed version.txt
var version string

func TestEmbedValue(t *testing.T) {
	fmt.Println(version)
}

// use "[]byte" for load picture, music, or video
//go:embed logo.svg
var logo []byte

func TestByte(t *testing.T) {
	err := ioutil.WriteFile("new_logo.svg", logo, fs.ModePerm) // create the file using the logo variable
	if err != nil {
		panic(err)
	}
}

// Embed multiple file
// if you want to embed multiple file, you have to use embed.FS as variable type
//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/c.txt
var files embed.FS

func TestMultipleFile(t *testing.T) {
	a, err := files.ReadFile("files/a.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(a))

	b, err := files.ReadFile("files/b.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	c, err := files.ReadFile("files/c.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(c))
}

// Path Matcher
// https://pkg.go.dev/path#Match
// to get all the files with the same pattern
// "go:embed files/*.txt" get all the file in files forlder with extention ".txt"

//go:embed files/*.txt
var path embed.FS

func TestPathMatcher(t *testing.T) {
	dir, err := path.ReadDir("files")
	if err != nil {
		panic(err)
	}
	for _, entry := range dir {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			content, err := path.ReadFile("files/" + entry.Name())
			if err != nil {
				panic(err)
			}
			fmt.Println("Content:", string(content))
		}
	}
}
