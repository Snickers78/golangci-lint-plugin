package rules

import (
	"fmt"
	"go/ast"
	"go/token"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

var LogAnalyzer = &analysis.Analyzer{
	Name: "logAnalyzer",
	Doc:  "Checks log messages formatting and content rules",
	Run:  run,
}

func run(pass *analysis.Pass) (any, error) {
	cfg := GetConfig()
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

				if cfg.CheckLowercase && isLowercase(msg) {
					fixed := toLowercaseFirst(msg)
					newLiteral := fmt.Sprintf("%q", fixed)
					pass.Report(analysis.Diagnostic{
						Pos:     str.Pos(),
						End:     str.End(),
						Message: "first letter in log message must be in lowercase",
						SuggestedFixes: []analysis.SuggestedFix{
							{
								Message: "make first letter lowercase",
								TextEdits: []analysis.TextEdit{
									{
										Pos:     str.Pos(),
										End:     str.End(),
										NewText: []byte(newLiteral),
									},
								},
							},
						},
					})
				}

				if cfg.CheckEnglishOnly && containsNonEnglishLetters(msg) {
					pass.Report(analysis.Diagnostic{
						Pos:     str.Pos(),
						End:     str.End(),
						Message: "log message must contain only english letters",
					})
				}

				if cfg.CheckSpecialSymbols && containsSpecialSymbolsOrEmoji(msg) {
					pass.Report(analysis.Diagnostic{
						Pos:     str.Pos(),
						End:     str.End(),
						Message: "log message must not contain special characters or emoji",
					})
				}

				if sensitive, reason := containsSensitiveData(msg); sensitive && cfg.CheckSensitiveData {
					pass.Report(analysis.Diagnostic{
						Pos:     str.Pos(),
						End:     str.End(),
						Message: fmt.Sprintf("log message appears to contain sensitive data: %s", reason),
					})
				}
			}

			if sensitiveVar, name, ident := containsSensitiveIdentifier(arg); sensitiveVar && cfg.CheckSensitiveData {
				pass.Report(analysis.Diagnostic{
					Pos:     ident.Pos(),
					End:     ident.End(),
					Message: fmt.Sprintf("log message appears to log sensitive variable %s", name),
					SuggestedFixes: []analysis.SuggestedFix{
						{
							Message: "mask sensitive variable",
							TextEdits: []analysis.TextEdit{
								{
									Pos:     ident.Pos(),
									End:     ident.End(),
									NewText: []byte(`"***"`),
								},
							},
						},
					},
				})
			}

			return true
		})
	}
	return nil, nil
}

func toLowercaseFirst(s string) string {
	runes := []rune(s)
	if len(runes) == 0 {
		return s
	}
	runes[0] = unicode.ToLower(runes[0])
	return string(runes)
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
