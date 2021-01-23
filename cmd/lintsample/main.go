package main

import (
	"golang.org/x/tools/go/analysis/unitchecker"
	"lintsample"
)

func main() { unitchecker.Main(lintsample.Analyzer) }
