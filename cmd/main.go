package main

import (
	"github.com/snickers78/golangci-lint-plugin/rules"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(rules.LogAnalyzer)
}
