package lintsample

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "lintsample is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "lintsample",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.ForStmt)(nil),
		(*ast.RangeStmt)(nil),
		(*ast.Ident)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.RangeStmt:
			checkRangeStmt(pass, n)
		}
	})

	return nil, nil
}

func checkRangeStmt(pass *analysis.Pass, rangeStmt *ast.RangeStmt) {
	ident, ok := rangeStmt.Key.(*ast.Ident)
	if !ok {
		return
	}
	assignStmt, ok := ident.Obj.Decl.(*ast.AssignStmt)
	if !ok {
		return
	}
	iterators := getIteratorIdent(assignStmt)
	if len(iterators) < 2 {
		return
	}
	reportWhenUsingSecondValue(pass, rangeStmt.Body, iterators[1])
}

func getIteratorIdent(stmt *ast.AssignStmt) []*ast.Ident {
	if len(stmt.Lhs) == 0 {
		return nil
	}

	var iterators []*ast.Ident
	for _, expr := range stmt.Lhs {
		switch expr := expr.(type) {
		case *ast.Ident:
			if expr.Obj.Kind == ast.Var {
				iterators = append(iterators, expr)
			}
		}
	}

	return iterators
}

func reportWhenUsingSecondValue(pass *analysis.Pass, stmt ast.Stmt, iterator *ast.Ident) {
	ast.Inspect(stmt, func(n ast.Node) bool {
		if n == nil {
			return false
		}

		switch n := n.(type) {
		case *ast.UnaryExpr:
			_, ok := n.X.(*ast.Ident)
			if !ok {
				return true
			}
			pass.Reportf(n.Pos(), "should use identifier of slice or array")
		case *ast.Ident:
			if n.Obj == iterator.Obj {
				pass.Reportf(n.Pos(), "should use identifier of slice or array")
			}
		}

		return true
	})
}
