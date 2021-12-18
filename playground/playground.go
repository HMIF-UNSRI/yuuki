package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	basename := "hello.blah"
	fmt.Println(filepath.Ext(basename))
}
