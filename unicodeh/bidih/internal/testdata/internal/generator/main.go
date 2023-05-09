package main

import (
	"eval_helper/unicodeh/bidih/internal/testdata/internal/generator/internal/lib"
	"os"
)

func main() {
	const usage = "Bad usage. Usage: \"maketables path-to-ucd-directory\""
	if len(os.Args) != 2 {
		panic(usage)
	}
	srcDir := os.Args[1]

	// TODO remove old data files
	lib.GenerateBidiTests(srcDir)
	lib.GenerateBidiCharacterTests(srcDir)
	lib.GenerateRuneSamples()
}
