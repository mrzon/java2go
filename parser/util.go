package parser

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"
)

func makeVar(name string, texpr ast.Expr) *ast.Field {
	return makeFieldWithTag(name, "", texpr)
}

func makeField(name string, texpr ast.Expr) *ast.Field {
	return makeFieldWithTag(strings.Title(name), fmt.Sprintf("`json:\"%s\"`", name), texpr)
}

func makeFieldWithTag(name string, tag string, texpr ast.Expr) *ast.Field {
	fname := make([]*ast.Ident, 0)
	if name != "" {
		fname = append(fname, ast.NewIdent(name))
	}
	var tagVal *ast.BasicLit
	if tag != "" {
		tagVal = &ast.BasicLit{
			ValuePos: token.Pos(len(fname)),
			Kind:     token.STRING,
			Value:    tag,
		}
	}
	return &ast.Field{Names: fname, Type: texpr, Tag: tagVal}
}
