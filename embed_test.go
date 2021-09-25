package learn_go_embed

import (
	_ "embed"
	"fmt"
	"testing"
)

// "go:embed [name of the file]" to load/embed file in golang
// load it into a variable outside function
//go:embed version.txt
var version string

func TestEmbedValue(t *testing.T) {
	fmt.Println(version)
}
