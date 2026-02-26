package main

import (
	"github.com/snickers78/golangci-lint-plugin/rules"
	"golang.org/x/tools/go/analysis"
)

func New(conf any) ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{rules.LogAnalyzer}, nil
}
