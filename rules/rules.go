package rules

import (
	"go/ast"
	"go/token"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

var LogAnalyzer = &analysis.Analyzer{
	Name: "lowercase",
	Doc:  "Checks log messages formatting and content rules",
	Run:  run,
}

func run(pass *analysis.Pass) (any, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			node, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}
			if len(node.Args) == 0 || !isLogFunction(node) {
				return true
			}

			arg := node.Args[0]

			if str, ok := arg.(*ast.BasicLit); ok && str.Kind == token.STRING {
				msg := str.Value[1 : len(str.Value)-1]

				if isLowercase(msg) {
					pass.Reportf(node.Pos(), "first letter in log message must be in lowercase")
				}

				if containsNonEnglishLetters(msg) {
					pass.Reportf(node.Pos(), "log message must contain only english letters")
				}

				if containsSpecialSymbolsOrEmoji(msg) {
					pass.Reportf(node.Pos(), "log message must not contain special characters or emoji")
				}

				if sensitive, reason := containsSensitiveData(msg); sensitive {
					pass.Reportf(node.Pos(), "log message appears to contain sensitive data: %s", reason)
				}
			}

			if sensitiveVar, name := containsSensitiveIdentifier(arg); sensitiveVar {
				pass.Reportf(node.Pos(), "log message appears to log sensitive variable %s", name)
			}

			return true
		})
	}
	return nil, nil
}

func isLowercase(s string) bool {
	runes := []rune(s)
	if len(runes) > 0 && unicode.IsUpper(runes[0]) {
		return true
	}

	return false
}

func isLogFunction(call *ast.CallExpr) bool {
	if fun, ok := call.Fun.(*ast.SelectorExpr); ok {
		switch fun.Sel.Name {
		case "Info", "Error", "Warn", "Println", "Debug":
			return true
		}
	}
	return false
}
