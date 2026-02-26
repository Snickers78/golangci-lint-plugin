package rules

import (
	"go/ast"
	"strings"
)

func containsSensitiveIdentifier(expr ast.Expr) (bool, string) {
	var found string

	ast.Inspect(expr, func(n ast.Node) bool {
		ident, ok := n.(*ast.Ident)
		if !ok {
			return true
		}

		name := strings.ToLower(ident.Name)
		for _, sub := range []string{"token", "password", "secret", "apikey"} {
			if strings.Contains(name, sub) {
				found = ident.Name
				return false
			}
		}

		return true
	})

	if found != "" {
		return true, found
	}
	return false, ""
}
