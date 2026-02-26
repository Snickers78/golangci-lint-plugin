package rules

import (
	"go/ast"
	"strings"
)

func containsSensitiveIdentifier(expr ast.Expr) (bool, string, *ast.Ident) {
	var found *ast.Ident

	ast.Inspect(expr, func(n ast.Node) bool {
		ident, ok := n.(*ast.Ident)
		if !ok {
			return true
		}

		name := strings.ToLower(ident.Name)
		cfg := GetConfig()
		base := []string{"token", "password", "secret", "apikey"}
		for _, sub := range append(base, cfg.CustomSensitiveSubstrs...) {
			if strings.Contains(name, sub) {
				found = ident
				return false
			}
		}

		return true
	})

	if found != nil {
		return true, found.Name, found
	}
	return false, "", nil
}
