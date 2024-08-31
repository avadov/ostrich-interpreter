package evaluator

import (
	"ostrich-interpreter/ast"
	"ostrich-interpreter/object"
)

func quote(node ast.Node) object.Object {
	return &object.Quote{Node: node}
}
