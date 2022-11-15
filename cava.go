package main

import (
	"Cava/util"
	"flag"
)

func main() {
	compileLevel := flag.Int("l", 0, "set compiler optimization level")
	debug := flag.Bool("d", false, "use debug")
	remainArgs := flag.Args()
	flag.Parse()
	sourceFile := remainArgs[0]
	util.ReadBuildFile(sourceFile)
}
